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

func TestCreate(t *testing.T) {
	chapter01.CreateSingle()
}

func TestCreateMany(t *testing.T) {
	chapter01.CrateMany()
}

func TestSelect(t *testing.T) {
	chapter01.Select()
}

// 测试基于主键ID来进行数据查询
func TestSelectById(t *testing.T) {
	chapter01.SelectById()
}

// 测试查询所有数据
func TestSelectAll(t *testing.T) {
	chapter01.SelectAll()
}

// 测试string 条件查询
func TestSelectByConditionString(t *testing.T) {
	chapter01.SelectByConditionString()
}

// 查询结构体进行查询
func TestSelectByStruct(t *testing.T) {
	chapter01.SelectByStruct()
}

// 内联查询
func TestSelect4(t *testing.T) {
	chapter01.Select4()
}

// Not查询
func TestSelectNot(t *testing.T) {
	chapter01.SelectNot()
}

// or查询
func TestSelectOr(t *testing.T) {
	chapter01.SelectOr()
}

// 查询指定字段
func TestSelectFields(t *testing.T) {
	chapter01.SelectFields()
}

// 指定查询的排序字段及方式
func TestSelectOrder(t *testing.T) {
	chapter01.SelectOrder()
}

// 指定查询的排序字段及方式
func TestSelectLimit(t *testing.T) {
	chapter01.SelectLimit()
}

// group及having用法
func TestSelectGroup(t *testing.T) {
	chapter01.SelectGroup()
}

// Distinct
func TestSelectDistinct(t *testing.T) {
	chapter01.SelectDistinct()
}

// save
func TestSave(t *testing.T) {
	chapter01.Save()
}

// 测试更新单个列
func TestUpdateSingleCol(t *testing.T) {
	chapter01.UpdateSingleCol()
}

// 测试更新多个列
func TestUpdateMany(t *testing.T) {
	chapter01.UpdateMany()
}

func TestUpdateBatch(t *testing.T) {
	chapter01.UpdateBatch()
}
