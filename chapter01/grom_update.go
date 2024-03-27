package chapter01

import (
	"fmt"
	"time"
)

func Save() {
	var user User
	Db.First(&user)

	user.Name = "jinzhu 2"
	user.Age = 100
	// UPDATE `users` SET `name`='jinzhu 2',`email`=NULL,
	//`age`=100,`birthday`='2024-03-21 16:52:50.665',`member_number`=NULL,
	//`activated_at`=NULL,`created_at`='2024-03-21 16:52:50.69',
	//`updated_at`='2024-03-25 16:25:08.448' WHERE `id` = 1
	Db.Save(&user)
	// INSERT INTO `users` (`name`,`email`,`age`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`)
	// VALUES ('jinzhu8',NULL,22,NULL,NULL,NULL,'2024-03-25 16:38:58.857','2024-03-25 16:38:58.857')
	Db.Save(&User{
		Name:      "jinzhu8",
		Age:       22,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	// UPDATE `users` SET `name`='jinzhu',`email`=NULL,`age`=100,`birthday`=NULL,`member_number`=NULL,
	//`activated_at`=NULL,`created_at`='2024-03-25 16:38:58.859',`updated_at`='2024-03-25 16:38:58.86' WHERE `id` = 1
	Db.Save(&User{ID: 1, Name: "jinzhu", Age: 100, CreatedAt: time.Now()})
}

// 更新单个列
func UpdateSingleCol() {
	var user User
	// 根据条件更新
	// UPDATE `users` SET `name`='hello',`updated_at`='2024-03-25 16:57:52.216' WHERE age = 18;
	Db.Model(&User{}).Where("age = ?", 18).Update("name", "hello")

	// User 的 ID 是 1
	// UPDATE `users` SET `name`='hello',`updated_at`='2024-03-25 16:57:52.227' WHERE `id` = 1
	Db.Model(&User{ID: 1}).Update("name", "hello")

	// 根据条件和 model 的值进行更新
	// UPDATE `users` SET `name`='hello',`updated_at`='2024-03-25 16:57:52.23' WHERE age = 100
	Db.Model(&user).Where("age = ?", 100).Update("name", "hello")

	// UPDATE `users` SET `name`='hello',`updated_at`='2024-03-25 17:02:04.975'
	// 无法执行更新，gorm默认就会阻止全局更新
	result := Db.Model(&User{}).Update("name", "hello")
	fmt.Println(result)

}

func UpdateMany() {
	var user = User{
		ID: 1,
	}

	// 根据 `struct` 更新属性，只会更新非零值的字段
	// UPDATE `users` SET `name`='hello',`updated_at`='2024-03-25 17:14:10.48' WHERE `id` = 1
	Db.Model(&user).Updates(User{Name: "hello", Age: 0})

	// 根据 `map` 更新属性
	// UPDATE `users` SET `age`=0,`name`='hello',`updated_at`='2024-03-25 17:14:10.494' WHERE `id` = 1
	Db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 0})

	// 指定更新name字段
	// UPDATE `users` SET `age`=18,`name`='hello2',`updated_at`='2024-03-25 17:33:33.932' WHERE `id` = 1
	Db.Model(&User{ID: 1}).Select("name").Updates(map[string]interface{}{"name": "hello2", "age": 18, "active": false})
}

// 批量更新操作
func UpdateBatch() {
	// UPDATE `users` SET `name`='hello',`age`=18,`updated_at`='2024-03-25 17:41:13.394' WHERE name = 'hello2'
	Db.Model(User{}).Where("name = ?", "hello2").Updates(User{
		Name: "hello",
		Age:  18,
	})
}
