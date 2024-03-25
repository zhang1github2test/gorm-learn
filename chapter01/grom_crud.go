package chapter01

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func CreateSingle() {
	// 这里需要替换成自己的mysql地址及账号等
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

func CrateMany() {
	// 这里需要替换成自己的mysql地址及账号等
	db, err := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	if err != nil {

	}

	var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	db.Create(&users)

	// 如果表不存在，则进行自动创建
	db.AutoMigrate(&User{})
	result := db.Create(&users) // 通过数据的指针来创建
	fmt.Println("影响的数据行数为:", result.RowsAffected)
}

func Select() {
	// 这里需要替换成自己的mysql地址及账号等
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	// 获取第一条记录（主键升序）
	var user User

	// 获取第一条记录（主键升序）
	// 相当于SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	result := db.First(&user)
	result = db.First(&user)
	fmt.Println(user)

	user = User{}
	// 获取一条记录，没有指定排序字段
	db.Take(&user)
	// SELECT * FROM users LIMIT 1;
	fmt.Println(user)

	user = User{}
	// 获取最后一条记录（主键降序）
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	db.Last(&user)
	user = User{}

	// result.RowsAffected  返回找到的记录数
	//  result.Error      returns error or nil
	fmt.Println(result.RowsAffected, result.Error)

	// 检查 ErrRecordNotFound 错误
	errors.Is(result.Error, gorm.ErrRecordNotFound)

	// works because model is specified using `db.Model()`
	result1 := map[string]interface{}{}
	db.Model(&User{}).First(&result1)
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

	// doesn't work
	// model value required
	result1 = map[string]interface{}{}
	db.Table("users").First(&result1)
	fmt.Println(result1)

	// works with Take
	// SELECT * FROM `users` LIMIT 1
	result1 = map[string]interface{}{}
	db.Table("users").Take(&result1)

}

func SelectById() {

	// 这里需要替换成自己的mysql地址及账号等
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	var users []User
	var user2 User
	// 根据主键查询索引

	// SELECT * FROM users WHERE id = 10;
	result := db.First(&user2, 1)

	fmt.Println(result)

	user2 = User{}
	// SELECT * FROM users WHERE id = 1;
	db.First(&user2, "1")

	user2 = User{}

	// SELECT * FROM `users` WHERE id = 1 ORDER BY `users`.`id` LIMIT 1
	db.First(&user2, "id = ?", 1)

	// SELECT * FROM users WHERE id IN (1,2,3);
	db.Find(&users, []int{1, 2, 3})

	var user = User{ID: 1}
	// SELECT * FROM `users` WHERE `users`.`id` = 1 ORDER BY `users`.`id` LIMIT 1
	db.First(&user)
}

func SelectAll() {
	var users []User
	// 这里需要替换成自己的mysql地址及账号等
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	// SELECT * FROM `users`
	db.Find(&users)
	fmt.Println(users)
}

// string条件查询
func SelectByConditionString() {
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	var user User
	// SELECT * FROM `users` WHERE name = 'jinzhu' AND age = 18
	// 等于查询
	db.Where("name = ?", "jinzhu").Where("age = ?", 18).Find(&user)
	fmt.Println(user)

	var users []User
	// SELECT * FROM `users` WHERE name <> 'jinzhu'
	// 不等于查询
	db.Where("name <> ?", "jinzhu").Find(&users)
	fmt.Println(users)

	// SELECT * FROM `users` WHERE name IN ('jinzhu','jinzhu2')
	// in查询
	db.Where("name IN ?", []string{"jinzhu", "jinzhu2"}).Find(&users)
	fmt.Println(users)

	// SELECT * FROM `users` WHERE name like '%jinzhu%'
	// like模糊查询
	db.Where("name like ?", "%jinzhu%").Find(&users)
	fmt.Println(users)

	// SELECT * FROM `users` WHERE name = 'jinzhu' and age >= 10
	// and查询
	db.Where("name = ? and age >= ?", "jinzhu", 10).Find(&users)
	fmt.Println(users)
}

// 结构体及map查询
func SelectByStruct() {
	var user User
	var users []User
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	// Struct
	// SELECT * FROM `users` WHERE `users`.`name` = 'jinzhu' AND `users`.`age` = 18 ORDER BY `users`.`id` LIMIT 1;
	db.Where(&User{Name: "jinzhu", Age: 18}).First(&user)

	// Map
	// SELECT * FROM `users` WHERE `age` = 18 AND `name` = 'jinzhu'
	db.Where(map[string]interface{}{"name": "jinzhu", "age": 18}).Find(&users)

	// Slice of primary keys
	// SELECT * FROM `users` WHERE `users`.`id` IN (1,2,3)
	db.Where([]int64{1, 2, 3}).Find(&users)

	// 不会基于age添加构造查询条件
	// SELECT * FROM users WHERE name = "jinzhu";
	db.Where(&User{Name: "jinzhu", Age: 0}).Find(&users)

	// SELECT * FROM `users` WHERE `Age` = 0 AND `Name` = 'jinzhu'
	db.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users)

}

// 使用内联的方式进行条件查询
func Select4() {
	var user User
	var users []User
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	// Get by primary key if it were a non-integer type
	// SELECT * FROM users WHERE id = 1;
	db.First(&user, "id = ?", 1)

	// Plain SQL
	db.Find(&user, "name = ?", "jinzhu")
	// SELECT * FROM users WHERE name = "jinzhu";

	db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

	// Struct
	db.Find(&users, User{Age: 20})
	// SELECT * FROM users WHERE age = 20;

	// Map
	db.Find(&users, map[string]interface{}{"age": 20})
	// SELECT * FROM users WHERE age = 20;
}

// Not条件
func SelectNot() {
	var user User
	var users []User
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	db.Not("name = ?", "jinzhu").First(&user)
	// SELECT * FROM users WHERE NOT name = "jinzhu" ORDER BY id LIMIT 1;

	// Not In
	db.Not(map[string]interface{}{"name": []string{"jinzhu", "jinzhu 2"}}).Find(&users)
	// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");

	// Struct
	db.Not(User{Name: "jinzhu", Age: 18}).First(&user)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age <> 18 ORDER BY id LIMIT 1;

	// Not In slice of primary keys
	db.Not([]int64{1, 2, 3}).First(&user)
	// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;
}

// Or条件
func SelectOr() {
	var users []User
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")

	// SELECT * FROM `users` WHERE name = 'jinzhu' OR (name = 'jinzhu 2' and age = 18)
	db.Where("name = ?", "jinzhu").Or("name = ? and age = ?", "jinzhu 2", 18).Find(&users)

	// Struct
	// SELECT * FROM `users` WHERE name = 'jinzhu' OR (`users`.`name` = 'jinzhu 2' AND `users`.`age` = 18)
	db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2", Age: 18}).Find(&users)

	// Map
	// SELECT * FROM `users` WHERE name = 'jinzhu' OR (`age` = 18 AND `name` = 'jinzhu 2')
	db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2", "age": 18}).Find(&users)
}

// 查询指定的字段
func SelectFields() {
	var users []User
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")

	// SELECT `name`,`age` FROM `users`
	db.Select("name", "age").Find(&users)
}

// 指定数据排序
func SelectOrder() {
	var users []User
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	db.Order("age desc, name").Find(&users)
	// SELECT * FROM users ORDER BY age desc, name;

	// Multiple orders
	// SELECT * FROM users ORDER BY age desc, name;
	db.Order("age desc").Order("name").Find(&users)

}

// 分页查询
func SelectLimit() {
	var users []User
	var users1 []User
	var users2 []User
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	db.Limit(3).Find(&users)
	// SELECT * FROM users LIMIT 3;

	// Cancel limit condition with -1
	db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	// SELECT * FROM users LIMIT 10; (users1)
	// SELECT * FROM users; (users2)

	// SELECT * FROM `users` LIMIT 10 OFFSET 5
	// limit表示获取多少条数据  OFFSET表示跳过多少条数据
	db.Offset(1).Limit(10).Find(&users)

}

type result struct {
	Date  time.Time
	Total int
}

func SelectGroup() {
	var result result
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")

	db.Model(&User{}).Select("name, sum(age) as total").Where("name LIKE ?", "group%").Group("name").Order("name desc").Find(&result)

	db.Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "group").Find(&result)
}

func SelectDistinct() {
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	var result []map[string]interface{}
	db.Model(&User{}).Distinct("name", "age").Find(&result)
	fmt.Println(result)
}
