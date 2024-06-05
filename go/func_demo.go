package main

import "fmt"

func main() {
	result := add(3, 4)
	fmt.Println(result) // 输出：7

	quotient, remainder := divide(8, 2)
	fmt.Println("quotient:", quotient, "remainder:", remainder)

	s := sum(1, 2, 3, 4, 5, 6)

	fmt.Println("sum的和为:", s)
}

// 函数的基本定义
func add(a, b int) int {
	return a + b
}

// 多返回值的函数定义方法
func divide(a int, b int) (int, int) {

	if b == 0 {

		return 0, 1 // 返回商和错误码（这里用1表示除以零）

	}

	quotient := a / b

	remainder := a % b

	return quotient, remainder

}

// 可变参数函数的定义方法
func sum(numbers ...int) int {

	total := 0

	for _, num := range numbers {

		total += num

	}

	return total

}
