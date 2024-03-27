package chapter01

// 演示删除一条数据
func DeleteSingle() {
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	// 删除Id = 1的数据
	// DELETE FROM `users` WHERE `users`.`id` = 1
	db.Delete(&User{
		ID: 1,
	})
}

// 通过主键删除数据
func DeleteById() {
	var users []User
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	// DELETE FROM users WHERE id = 10;
	db.Delete(&User{}, 10)

	// DELETE FROM users WHERE id = 10;
	db.Delete(&User{}, "10")

	// DELETE FROM `users` WHERE `users`.`id` IN (1,2,3)
	db.Delete(&users, []int{1, 2, 3})
}

// 按照条件删除对应的
func DeleteBatch() {
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")

	// DELETE FROM `users` WHERE name = 'jinzhu'
	db.Delete(&User{}, "name = ?", "jinzhu")
}
