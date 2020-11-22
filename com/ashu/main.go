package main

import (
	"./localrequest"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 绑定端口，然后启动应用
	engine := gin.Default()
	engine.Any("/", webRoot)
	engine.GET("/txDetail/:txHash", getTxDetail)
	engine.Run(":9999")
}

func webRoot(context *gin.Context) {
	var ipAddr = context.ClientIP()
	context.String(http.StatusOK, "Hello "+ipAddr)
}

func getTxDetail(context *gin.Context) {
	txHash := context.Param("txHash")
	var result = localrequest.Parse(txHash)
	context.String(http.StatusOK, result)
	return
}
