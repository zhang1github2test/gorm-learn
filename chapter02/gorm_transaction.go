package chapter02

import (
	"errors"
	"fmt"
	"go-orm-learn/chapter01"
	"gorm.io/gorm"
)

func ContextDemo() {
	err := chapter01.Db.Transaction(multiOpertaion)
	if err != nil {
		fmt.Println("执行事务出错")
		errors.New("执行事务出错")
	}

}

func multiOpertaion(tx *gorm.DB) error {
	tx.Create(&chapter01.User{
		Name: "zhshl",
	})
	// return errors.New("返回错误！数据库中将不会生成新的数据！")
	return nil
}
