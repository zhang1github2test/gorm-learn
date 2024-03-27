package chapter01

// 演示删除一条数据
func DeleteSingle() {
	// 删除Id = 1的数据
	// DELETE FROM `users` WHERE `users`.`id` = 1
	Db.Delete(&User{
		ID: 1,
	})
}

// 通过主键删除数据
func DeleteById() {
	var users []User
	// DELETE FROM users WHERE id = 10;
	Db.Delete(&User{}, 10)

	// DELETE FROM users WHERE id = 10;
	Db.Delete(&User{}, "10")

	// DELETE FROM `users` WHERE `users`.`id` IN (1,2,3)
	Db.Delete(&users, []int{1, 2, 3})
}

// 按照条件删除对应的
func DeleteBatch() {
	// DELETE FROM `users` WHERE name = 'jinzhu'
	Db.Delete(&User{}, "name = ?", "jinzhu")
}
