package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64 // 零值
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

func (u User) TableName() string {
	// 实现TableName函数，自定义表名称
	return "user"
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_one?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("connect database success", db)

	// 创建表
	err = db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err)
	}

	err = db.Table("xiaowangzi").AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err)
	}

}
