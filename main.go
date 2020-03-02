package main

import (
	"blog/init"
	"blog/router"
	"fmt"
)

func main() {
	// 初始化配置
	globalInit.Now()

	// 装在路由
	r:=router.InitRouter()

	// 运行
	if err:=r.Run(":8888");err!=nil{
		fmt.Println("错误：",err.Error())
	}
}
