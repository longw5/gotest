package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   uint64
	Name string
}

func main() {
	users1 := []User{{ID: 1, Name: "张三"}, {ID: 456, Name: "李四"}}
	users2 := []User{{ID: 2, Name: "李四"}, {ID: 456, Name: "李四"}}
	users3 := []User{{ID: 3, Name: "王五"}, {ID: 456, Name: "李四"}}
	users4 := []User{{ID: 4, Name: "马六"}, {ID: 456, Name: "李四"}}
	users5 := []User{{ID: 5, Name: "田七"}, {ID: 456, Name: "李四"}}
	users6 := []User{{ID: 6, Name: "孙八"}, {ID: 456, Name: "李四"}}

	r := gin.Default()
	r.GET("/users1", func(c *gin.Context) {
		c.JSON(200, users1)
	})

	r.GET("/users2", func(c *gin.Context) {
		c.JSON(200, users2)
	})

	r.GET("/users3", func(c *gin.Context) {
		c.JSON(200, users3)
	})

	r.GET("/users4", func(c *gin.Context) {
		c.JSON(200, users4)
	})

	r.GET("/users5", func(c *gin.Context) {
		c.JSON(200, users5)
	})

	r.GET("/users6", func(c *gin.Context) {
		c.JSON(200, users6)
	})

	r.POST("/users", func(context *gin.Context) {
		//创建一个用户
	})
	r.DELETE("/usrs/123", func(context *gin.Context) {
		//删除ID为123的用户
	})
	r.PUT("/usrs/123", func(context *gin.Context) {
		//更新ID为123的用户
	})

	r.PATCH("/usrs/123", func(context *gin.Context) {
		//更新ID为123用户的部分信息
	})

	r.GET("/users/op1/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "The user id is  %s", id)
	})

	r.GET("/users/op2/*name", func(c *gin.Context) {
		id := c.Param("name")
		c.String(200, "The user id is  %s", id)
	})

	r.GET("/query", func(c *gin.Context) {

		fmt.Println(c.Query("aa"))
		c.String(200, c.Query("aa"))
	})

	r.Run(":8080")
}
