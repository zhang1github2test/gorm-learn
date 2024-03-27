package chapter01

import "fmt"

// 通过原生sql操作数据
func SelectBySql() {
	var user User
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	db.Raw("select * from users;").Scan(&user)
	fmt.Println(user)
}
