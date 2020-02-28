package main

import (
	"fmt"
	"time"
)

var male chan string
var female chan string

func main() {
	/*
		select {
		case <- chan1 :
			// 如果 chan1 成功读到数据，则进行该 case 处理语句
		case chan2:
			// 如果成功向 chan2写入数据，则进行该 case 处理语句
		default:
			// 如果上面都没有成功，则进入 default 处理流程
		}
		PS：select 中当执行任意 case 处理语句后，不再执行其他 case 处理语句

		// 真正的程序会利用 for 循环，使 select 一直轮询检查
		for {
			select {

			}
		}
	*/

	ch1 := make(chan string)
	ch2 := make(chan string)

	go test1(ch1)
	go test2(ch2)
	time.Sleep(10 * time.Millisecond)

	select {
	case <-ch1:
		fmt.Println("come to read ch1!")
	case <-ch2:
		fmt.Println("come to read ch2!")
	default:
		fmt.Println("come to default!")
	}
}

func test1(ch chan string) {
	ch <- "男"
}

func test2(ch chan string) {
	ch <- "女"
}
