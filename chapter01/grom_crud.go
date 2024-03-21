package chapter01

import (
	"fmt"
	"time"
)

func CreateSingle() {
	db, err := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	if err != nil {

	}
	birthTime := time.Now()
	user := User{Name: "Jinzhu", Age: 18, Birthday: &birthTime}

	// 如果表不存在，则进行自动创建
	db.AutoMigrate(&User{})
	result := db.Create(&user) // 通过数据的指针来创建
	fmt.Println("影响的数据行数为:", result.RowsAffected)
}
