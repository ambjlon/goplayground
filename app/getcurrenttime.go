package app

import (
	"fmt"
	//"testing"
	"time"
)

//测试获得当前时间, 并以各种形式显示
func TestGetCurrentTime() {
	t1 := time.Now().Unix()
	fmt.Println(t1)
}
