package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := InitRouter()             // 注册路由
	r.Run(fmt.Sprintf(":%d", 80)) // 启动服务
}

func InitRouter() *gin.Engine {
	var r = gin.Default()
	r.GET("/test", func(context *gin.Context) {
		context.Writer.Write([]byte("hi"))
	})
	//r.Any("/*any", rpc.ReverseProxy)
	return r
}
