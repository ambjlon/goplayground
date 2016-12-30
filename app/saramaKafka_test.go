package app

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/bitly/go-simplejson"
	"github.com/wvanbergen/kafka/consumergroup"
	"testing"
	"time"
)

//和xiaowan一起测试该kafka客户端对offset的处理机制.
func TestConsumerOffset(t *testing.T) {
	group := "***_collect_group1" // group可以有客户端自主定义
	topics := []string{"***_collect_topic"}
	zkhosts := []string{"*.*.*.*:2181"}

	config := consumergroup.NewConfig()
	config.Zookeeper.Chroot = "/kafka/Test01"       // zk node??
	config.Offsets.Initial = sarama.OffsetOldest    // 只在一个group第一次连接kafka的时候起作用, 或者在offset不存在的时候存在. new是从最新的下一个消息开始读, old是从最老的消息开始读.
	config.Offsets.CommitInterval = 2 * time.Second // 每个两秒提交一下消息的commit给kafka, 也就是相应的offset做更改.

	consumer, err := consumergroup.JoinConsumerGroup(group, topics, zkhosts, config)
	if err != nil {
		fmt.Println("connect server error!")
		return
	}

	i := 0
	for {
		if i == 3 {
			time.Sleep(10 * time.Second)
			return
		}
		select {
		case msg := <-consumer.Messages():
			fmt.Println(string(msg.Value)+"\t"+"offset", msg.Offset)
			consumer.CommitUpto(msg) //执行commit
			i = i + 1
		}
	}
}

//给qa的简单的测试小工. 用来生产消息.
func TestSaramaKafkaProducer(ts *testing.T) {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	brokers := []string{"*.*.*.*:9092"}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		// Should not reach here
		fmt.Printf("connect to zk error![error=%v]\n", err)
		return
	}

	defer func() {
		if err := producer.Close(); err != nil {
			// Should not reach here
			fmt.Println("close connect error!")
		}
	}()

	topic := "***_topic"
	var t time.Time
	var secs int64
	var tmpJSON *simplejson.Json
	var tmpMsg []byte
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(string("")),
	}

	t = time.Now()
	secs = t.Unix()
	tmpJSON, _ = simplejson.NewJson([]byte(MSG_A))
	tmpJSON.Set("createtime", secs)
	tmpMsg, _ = tmpJSON.MarshalJSON()
	msg = &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(string(tmpMsg)),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Printf("send kafka msg error![error=%v]", err)
		return
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	time.Sleep(80 * time.Second)
	t = time.Now()
	secs = t.Unix()
	tmpJSON, _ = simplejson.NewJson([]byte(MSG_B))
	tmpJSON.Set("createtime", secs)
	tmpMsg, _ = tmpJSON.MarshalJSON()
	msg = &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(string(tmpMsg)),
	}
	partition, offset, err = producer.SendMessage(msg)

	if err != nil {
		fmt.Printf("send kafka msg error![error=%v]", err)
		return
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
}

const (
	MSG_A = `{
    "type": "AC",
    "createtime": 1482727714,
    "op": 600}`
	MSG_B = `{
    "type": "BC",
    "createtime": 1482727714,
    "op": 600}`
)
