package main

import "go-orm-learn/chapter01"

var (
	JK, _ = chapter01.GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
)
