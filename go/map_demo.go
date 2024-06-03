package main

import "fmt"

func main() {
	// 申明指针
	var aPtr *int

	var num = 4
	// 进行指针赋值
	aPtr = &num
	// 获取指针指向的值 输出4
	fmt.Println("通过指针获取到的num变量的值为：", *aPtr)
	fmt.Println("指针的地址值为:", &aPtr)

	// 通过指针修改变量的值
	*aPtr = 20
	// 此时值会被修改为20
	fmt.Println("通过指针修改num变量的值为：", *aPtr)

	modifyParam(aPtr)
	// 此时值会被修改为30
	fmt.Println("通过函数传递指针参数修改后，num变量的值为：", *aPtr)

}
func modifyParam(n *int) {
	*n = 30
}

// 交换a,b的值
// 指针作为函数参数传递，以修改函数外部的变量
func swap(a, b *int) {
	*a, *b = *b, *a
}
