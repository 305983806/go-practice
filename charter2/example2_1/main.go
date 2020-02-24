package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("------- sample1 ------")
	fmt.Println("Start...")
	go test_Routine()


	fmt.Println("------- sample2 ------")
	for i := 1; i < 10; i++ {
		go add(i, i)
	}

	// 加入阻塞，让带有 go 关键字的协程获得控制权
	time.Sleep(10 * time.Millisecond)

	fmt.Println("stop...")
}

func test_Routine()  {
	fmt.Println("This is a routine!!!")
}

func add(a, b int) {
	c := a + b
	fmt.Println(c)
}