package chapter01

import (
	"fmt"
	"testing"
)

func TestMysqlConnection(t *testing.T) {
	db, err := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	if err != nil {
		t.Errorf("err is not nil")
	}
	sqldb, _ := db.DB()

	fmt.Println("已经建立的链接数", sqldb.Stats().OpenConnections)
}

func TestCreate(t *testing.T) {
	CreateSingle()
}

func TestCreateMany(t *testing.T) {
	CrateMany()
}

func TestSelect(t *testing.T) {
	Select()
}

// 测试基于主键ID来进行数据查询
func TestSelectById(t *testing.T) {
	SelectById()
}

// 测试查询所有数据
func TestSelectAll(t *testing.T) {
	SelectAll()
}

// 测试string 条件查询
func TestSelectByConditionString(t *testing.T) {
	SelectByConditionString()
}

// 查询结构体进行查询
func TestSelectByStruct(t *testing.T) {
	SelectByStruct()
}

// 内联查询
func TestSelect4(t *testing.T) {
	Select4()
}

// Not查询
func TestSelectNot(t *testing.T) {
	SelectNot()
}

// or查询
func TestSelectOr(t *testing.T) {
	SelectOr()
}

// 查询指定字段
func TestSelectFields(t *testing.T) {
	SelectFields()
}

// 指定查询的排序字段及方式
func TestSelectOrder(t *testing.T) {
	SelectOrder()
}

// 指定查询的排序字段及方式
func TestSelectLimit(t *testing.T) {
	SelectLimit()
}

// group及having用法
func TestSelectGroup(t *testing.T) {
	SelectGroup()
}

// Distinct
func TestSelectDistinct(t *testing.T) {
	SelectDistinct()
}

// save
func TestSave(t *testing.T) {
	Save()
}

// 测试更新单个列
func TestUpdateSingleCol(t *testing.T) {
	UpdateSingleCol()
}

// 测试更新多个列
func TestUpdateMany(t *testing.T) {
	UpdateMany()
}

func TestUpdateBatch(t *testing.T) {
	UpdateBatch()
}
