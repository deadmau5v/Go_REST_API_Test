package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var books = make([]Book, 0, 512)
var bookNames = make(map[string]bool)
var bookIds = make(map[string]bool)

type Book struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func GetBook(ctx *gin.Context) {
	var data = make(map[string]interface{})
	for i, book := range books {
		data[strconv.Itoa(i)] = book
	}
	ctx.JSON(http.StatusOK, data)
}

// RemoveBook 删除书
func RemoveBook(ctx *gin.Context) {
	var newBooks []Book
	var data = make(map[string]interface{})
	Id := ctx.Query("id")

	// 如果ID为空
	if Id == "" {
		data["msg"] = "Id 不能为空!"
		data["status"] = "-1"
		ctx.JSON(http.StatusOK, data)
		return
	} else if !bookIds[Id] {
		data["msg"] = "Id 不存在!"
		data["status"] = "-2"
		ctx.JSON(http.StatusOK, data)
		return
	}

	for _, _book := range books {
		if _book.Id != Id {
			newBooks = append(newBooks, _book)
		} else {
			delete(bookNames, _book.Name)
			delete(bookIds, _book.Id)
		}
	}
	books = newBooks
	data["msg"] = "ok"
	data["status"] = "0"
	ctx.JSON(http.StatusOK, data)
	return
}

func AddBook(ctx *gin.Context) {
	var book Book
	var data = make(map[string]interface{})
	book.Name = ctx.PostForm("name")
	book.Id = ctx.PostForm("id")

	if book.Name == "" || book.Id == "" {
		data["msg"] = "书名和ID不能为空!"
		data["status"] = "-1"
	} else if bookIds[book.Id] {
		data["msg"] = "ID已存在!"
		data["status"] = "-2"
	} else if bookNames[book.Name] {
		data["msg"] = "书名已存在!"
		data["status"] = "-3"
	} else {
		data["msg"] = "ok"
		data["status"] = "0"
		bookNames[book.Name] = true
		bookIds[book.Id] = true
		books = append(books, book)
	}

	ctx.JSON(http.StatusOK, data)
}

func AlterBook(ctx *gin.Context) {
	var newBooks []Book
	var data = make(map[string]interface{})
	Id := ctx.PostForm("id")
	Name := ctx.PostForm("name")
	// 如果ID为空
	if Id == "" {
		data["msg"] = "Id 不能为空!"
		data["status"] = "-1"
		ctx.JSON(http.StatusOK, data)
		return
	} else if !bookIds[Id] {
		data["msg"] = "Id 不存在!"
		data["status"] = "-2"
		ctx.JSON(http.StatusOK, data)
		return
	} else if bookNames[Name] {
		data["msg"] = "Name 已存在!"
		data["status"] = "-3"
		ctx.JSON(http.StatusOK, data)
		return
	}

	for _, _book := range books {
		if _book.Id != Id {
			newBooks = append(newBooks, _book)
		} else {
			delete(bookNames, _book.Name)
			bookNames[Name] = true

			var book Book
			book.Name = Name
			book.Id = Id
			newBooks = append(newBooks, book)

		}
	}
	books = newBooks
	data["msg"] = "ok"
	data["status"] = "0"
	ctx.JSON(http.StatusOK, data)
	return
}

func BookFlushall(ctx *gin.Context) {
	books = make([]Book, 0, 512)
	bookNames = make(map[string]bool)
	bookIds = make(map[string]bool)
	data := make(map[string]interface{})
	data["msg"] = "ok"
	data["status"] = "0"
	ctx.JSON(http.StatusOK, data)
}
