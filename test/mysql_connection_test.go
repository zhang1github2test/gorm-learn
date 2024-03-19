package test

import (
	"fmt"
	"go-orm-learn/chapter01"
	"testing"
)

func TestMysqlConnection(t *testing.T) {
	db, err := chapter01.GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	if err != nil {
		t.Errorf("err is not nil")
	}
	sqldb, _ := db.DB()

	fmt.Println("已经建立的链接数", sqldb.Stats().OpenConnections)
}
