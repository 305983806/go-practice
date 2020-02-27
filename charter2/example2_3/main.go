package main

import (
	"fmt"
	"time"
)

var ch chan int

func main() {
	/*
		1）由于 read() 中没有接收 ch，因此 ch所在的协程 在执行完 ch <- 10 后即刻进入阻塞，导致无法执行 fmt.Println("goroutine is running...")
		最终只输出：“running end!”，
		2）如果希望 “goroutine is running...” 也可以被正常输出，需要这样操作：
		  a) 声明 ch 时，增加缓存区参数1（或更大），此时即便没有 <-ch（接收ch），整个协程也不会被阻塞；
		3）当向 ch 多次赋值，依然无法 输出 “goroutine is running...”，缓冲区大小仅为1，需要加大缓冲区大小至大于赋值次数，才不会阻塞
	*/

	ch := make(chan int)
	go write(ch)
	go read(ch)
	time.Sleep(10 * time.Millisecond)
	fmt.Println("running end!")
	time.Sleep(10 * time.Millisecond)

}

func read(ch chan int) {
	// value := <-ch
	// fmt.Println("value: ", strconv.Itoa(value))
	// <-ch
}

func write(ch chan int) {
	ch <- 10
	fmt.Println("goroutine is running...")
}
