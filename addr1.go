package main

import "fmt"

func main() {
	var a int = 10

	fmt.Printf("变量的地址: %x\n", &a)

	fmt.Printf("变量的地址: %d\n", a)

}
