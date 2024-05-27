package chapter02

import (
	"context"
	"errors"
	"fmt"
	. "go-orm-learn/chapter01"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// 在事务中执行一系列操作
func TransactionDemo() {
	ctx := context.WithValue(context.Background(), "a", "v")
	Db.WithContext(ctx)
	err := Db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&User{
			Name: "zhshl",
		})

		tx.Create(&User{
			Name: "zhshl2",
		})
		// return errors.New("返回错误！数据库中将不会生成新的数据！")
		// 不返回错误，那么将会
		return nil
	})
	if err != nil {
		fmt.Println("执行事务出错")
		errors.New("执行事务出错")
	}

}

// 执行嵌套事务
// user1 和user3对象会被成功创建，但是user2对象不会被创建,由于对象返回了一个错误，所以会被回滚掉
func TransactionDemo2() {
	user1 := User{
		Name: "zhangshl21",
	}
	user2 := User{
		Name: "zhangshl22",
	}

	user3 := User{
		Name: "zhangshl23",
	}
	Db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&user1)
		tx.Transaction(func(tx2 *gorm.DB) error {
			tx2.Create(&user2)
			return errors.New("模拟返回错误信息")
		})
		tx.Create(&user3)
		return nil
	})
}

// 手动事务演示
func TransactionDemo3() {
	user := User{
		Name: "manual transaction Rollback",
	}
	// 开始事务
	tx := Db.Begin()

	// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
	tx.Create(&user)

	// ...

	// 遇到错误时回滚事务
	tx.Rollback()

	tx = Db.Begin()
	user2 := User{
		Name: "manual transaction commit",
	}
	tx.Create(&user2)
	// 否则，提交事务
	tx.Commit()

}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
