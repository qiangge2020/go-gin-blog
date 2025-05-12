package main

import (
	"fmt"
	"github.com/qiangge2020/go-gin-blog/models"
	"github.com/qiangge2020/go-gin-blog/pkg/gredis"
	"github.com/qiangge2020/go-gin-blog/pkg/logging"
	"github.com/qiangge2020/go-gin-blog/pkg/setting"
	"github.com/qiangge2020/go-gin-blog/routers"
	"log"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	r := routers.InitRouter()
	err := r.Run(fmt.Sprintf(":%d", setting.ServerSetting.HttpPort))
	if err != nil {
		log.Printf("Server err: %v", err)
	}
	//读取请求头超时：服务器等待客户端发送请求头的最大时间
	//endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	////写入响应超时：服务器发送响应给客户端的最大时间
	//endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	////最大请求头字节数：限制HTTP请求头的大小
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	//server := endless.NewServer(endPoint, routers.InitRouter())
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}
