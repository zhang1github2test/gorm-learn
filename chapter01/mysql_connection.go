package chapter01

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetMysqlDb(account string, password string, host string, port int, dbname string) (*gorm.DB, error) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// 链接格式： [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", account, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
