package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("/book", GetBook)
	server.POST("/book", AddBook)
	server.PUT("/book", AlterBook)
	server.DELETE("/book", RemoveBook)
	server.GET("/book/flush", BookFlushall)

	_ = server.Run("127.0.0.1:8080")
}
