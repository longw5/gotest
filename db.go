package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo --> 数据表
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	//连接MySQL数据库
	db, err := gorm.Open("mysql", "root:root@(172.16.150.202)/godb?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 创建表、自动迁移(把结构体和数据表进行对应)
	db.AutoMigrate(&UserInfo{})

	// 创建数据行
	u1 := UserInfo{1, "Hina", "male", "read"}
	db.Create(&u1)

	// 查询
	var u UserInfo
	db.First(&u) // 查询表中第一条数据保存到u中
	fmt.Println(u)

	// 更新
	db.Model(&u).Update("hobby", "write")

	db.Delete(&u)
}
