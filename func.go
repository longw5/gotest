package main // 声明 main 包

// 导入内置 fmt 包
/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func main() { // main函数，程序执行入口

	println(max(1, 2))
}
