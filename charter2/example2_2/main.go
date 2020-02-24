package main

import (
	"fmt"
	"strconv"
)

func main() {
	// channel 作为协程的同步工具，实现协程间数据同步、通讯

	fmt.Println("------- sample1 ------")
	//var ch chan int
	//ch := make(chan int)
	ch := make(chan int)
	go write(ch)
	go read(ch)
	//time.Sleep(10 * time.Millisecond)
	fmt.Println("end of code...")

	fmt.Println("------- sample3 ------")
	chs := make([]chan int, 10)

	for i := 0; i < 10 ; i++ {
		chs[i] = make(chan int)
		go add(i, i, chs[i])
	}

	for _, v := range chs {
		<- v
	}
}

func read(ch chan int)  {
	value := <- ch
	fmt.Println("value: ", strconv.Itoa(value))
}

func write(ch chan int)  {
	ch <- 10
}

func add(a, b int, quit chan int) {
	c := a + b
	fmt.Println(c)

	quit <- 1
}