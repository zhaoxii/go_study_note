package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_one?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("connect database success")

	err = db.AutoMigrate(&UserInfo{})
	if err != nil {
		fmt.Println("auto migrate err:", err)
	}

	//u1 := UserInfo{ID: 1, Name: "zz", Gender: "n", Hobby: "aa"}
	//db.Create(&u1)

	var u UserInfo
	db.First(&u)
	fmt.Println("u:", u)
	db.Model(&u).Update("hobby", "哈哈哈")
	db.Delete(&u)
}
