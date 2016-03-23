package main

import (
	"fmt"
	"github.com/ambjlon/goplayground/app"
)

func main() {
	fmt.Println("my golang playground!")
	//app.TestViper()//暂时不可用, 丢了数据文件
	app.TestGetCurrentTime()
}
