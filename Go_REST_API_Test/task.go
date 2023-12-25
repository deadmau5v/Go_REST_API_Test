package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Tasks = make([]Task, 0, 512)
var TaskNames = make(map[string]bool)
var TaskIds = make(map[string]bool)

type Task struct {
	Name string `json:"name"`
	IsOk bool   `json:"isOk"`
	Id   string `json:"id"`
}

func GetTask(ctx *gin.Context) {
	var data = make(map[string]interface{})
	for i, Task := range Tasks {
		data[strconv.Itoa(i)] = Task
	}
	ctx.JSON(http.StatusOK, data)
}

// RemoveTask 删除Task
func RemoveTask(ctx *gin.Context) {
	var newTasks []Task = make([]Task, len(Tasks), len(Tasks)+1)
	var data = make(map[string]interface{})
	Id := ctx.Query("id")

	// 如果ID为空
	if Id == "" {
		data["msg"] = "Id 不能为空!"
		data["status"] = "-1"
		ctx.JSON(http.StatusBadRequest, data)
		return
	} else if !TaskIds[Id] {
		data["msg"] = "Id 不存在!"
		data["status"] = "-2"
		ctx.JSON(http.StatusBadRequest, data)
		return
	}

	for _, _Task := range Tasks {
		if _Task.Id != Id {
			newTasks = append(newTasks, _Task)
		} else {
			delete(TaskNames, _Task.Name)
			delete(TaskIds, _Task.Id)
		}
	}
	Tasks = newTasks
	data["msg"] = "ok"
	data["status"] = "0"
	ctx.JSON(http.StatusOK, data)
	return
}

func AddTask(ctx *gin.Context) {
	var Task Task
	var data = make(map[string]interface{})
	Task.Name = ctx.PostForm("name")
	Task.Id = ctx.PostForm("id")

	if Task.Name == "" || Task.Id == "" {
		data["msg"] = "Task名和ID不能为空!"
		data["status"] = "-1"
	} else if TaskIds[Task.Id] {
		data["msg"] = "ID已存在!"
		data["status"] = "-2"
	} else if TaskNames[Task.Name] {
		data["msg"] = "Task名已存在!"
		data["status"] = "-3"
	} else {
		data["msg"] = "ok"
		data["status"] = "0"
		TaskNames[Task.Name] = true
		TaskIds[Task.Id] = true
		Tasks = append(Tasks, Task)
		ctx.JSON(http.StatusOK, data)
		return
	}
	ctx.JSON(http.StatusBadRequest, data)
}

func AlterTask(ctx *gin.Context) {
	var newTasks []Task
	var data = make(map[string]interface{})
	Id := ctx.PostForm("id")
	Name := ctx.PostForm("name")
	// 如果ID为空
	if Id == "" {
		data["msg"] = "Id 不能为空!"
		data["status"] = "-1"
		ctx.JSON(http.StatusBadRequest, data)
		return
	} else if !TaskIds[Id] {
		data["msg"] = "Id 不存在!"
		data["status"] = "-2"
		ctx.JSON(http.StatusBadRequest, data)
		return
	} else if TaskNames[Name] {
		data["msg"] = "Name 已存在!"
		data["status"] = "-3"
		ctx.JSON(http.StatusBadRequest, data)
		return
	}

	for _, _Task := range Tasks {
		if _Task.Id != Id {
			newTasks = append(newTasks, _Task)
		} else {
			delete(TaskNames, _Task.Name)
			TaskNames[Name] = true

			var Task Task
			Task.Name = Name
			Task.Id = Id
			newTasks = append(newTasks, Task)

		}
	}
	Tasks = newTasks
	data["msg"] = "ok"
	data["status"] = "0"
	ctx.JSON(http.StatusOK, data)
	return
}

func TaskFlushall(ctx *gin.Context) {
	Tasks = make([]Task, 0, 512)
	TaskNames = make(map[string]bool)
	TaskIds = make(map[string]bool)
	data := make(map[string]interface{})
	data["msg"] = "ok"
	data["status"] = "0"
	ctx.JSON(http.StatusOK, data)
}
