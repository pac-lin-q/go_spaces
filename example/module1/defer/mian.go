package main

import (
	"fmt"
)

// defer进行压栈操作，先进后出，在函数结束后执行
func main() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer func() {
		fmt.Println("defer func is called")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("a panic is triggered")
}
