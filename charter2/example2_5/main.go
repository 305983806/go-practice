package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	select {
	case <-ch:
		// ch 没有被赋值，因此 case<-ch 中，ch无法接收到值，进入阻塞；
		fmt.Println("come to read ch!")
	case <-time.After(time.Second):
		// 当上面的都没有成功，且超时 1 秒后，则执行此代码逻辑
		fmt.Println("come to timeout!")
	}
}
