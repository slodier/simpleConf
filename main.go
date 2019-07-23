package main

import (
	"fmt"
	"simpleConf/conf"
)

func main() {
	conf.Init()		// 初始化

	fmt.Printf("%+v\n", conf.Tab)
	// &{Server:{Port:8010 LogModel:false} Database:{Host:localhost Port:3306 User:root Pw:123456qw} Image:{}}

}
