package chapter01

import "fmt"

// 通过原生sql操作数据
func SelectBySql() {
	var user User
	Db.Raw("select * from users;").Scan(&user)
	fmt.Println(user)
}
