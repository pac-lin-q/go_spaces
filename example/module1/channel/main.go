package main

import "fmt"

func main() {
	//for i := 0; i < 10; i++ {
	//	go fmt.Println(i)
	//}
	//time.Sleep(time.Second)
	ch := make(chan int)     //管道长度为0 需等管道两侧都就绪
	ch1 := make(chan int, 2) //新建管道长度为2，满后堵塞
	go func() {
		fmt.Println("hello from goroutine")
		ch <- 0 //数据写入Channel
	}()
	i := <-ch   //从Channel中取数据并赋值
	go func() { // 启动奕哥线程/协程进行处理
		fmt.Println("hello this is 2")
		ch1 <- 2
	}()
	j := <-ch1
	fmt.Println(i)
	fmt.Println(j)
}
