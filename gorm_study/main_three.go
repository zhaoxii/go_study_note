package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User2 struct {
	ID   int64
	Name string
	Age  int64
}

func (u User2) TableName() string {
	return "user"
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_two?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("connect database success", db)

	err = db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err)
	}
	db.Create(&User2{ID: 11, Name: "ada", Age: 121})

}
