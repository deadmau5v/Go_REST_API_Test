package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {

	server := gin.Default()

	server.GET("/task", GetTask)
	server.POST("/task", AddTask)
	server.PUT("/task", AlterTask)
	server.DELETE("/task", RemoveTask)
	server.GET("/task/flush", TaskFlushall)
	server.GET("/", index)

	println("Server Run in http://127.0.0.1:8080/ .")
	_ = server.Run("127.0.0.1:8080")

}

func index(ctx *gin.Context) {
	t, err := template.ParseFiles("./template/index.html")
	if err != nil {
		return
	}
	var Tasks []Task = make([]Task, len(Tasks), len(Tasks))

	for _, task := range Tasks {
		println(task.IsOk, task.Name, task.Id)
		Tasks = append(Tasks, task)
	}

	err = t.Execute(ctx.Writer, gin.H{
		"Tasks": Tasks,
	})

	if err != nil {
		println(err)
		return
	}
	return
}
