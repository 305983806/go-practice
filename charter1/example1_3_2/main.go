package main

import (
	"fmt"
)

/*
defer后面的函数在defer语句所在的函数执行结束的时候前才会被调用
作用1：类似于 JAVA try...catch...finally 中的 finally
注意：defer后面必须是函数调用语句或方法调用语句，不能是其他语句，否则编译器会出错。
*/
func main() {
	sample1()
	sample3()
}


func sample1() {
	for i := 0; i < 5; i++ {
		// defer 后的多个函数，也会根据声明顺序，倒序执行，输出 4 3 2 1 0
		defer fmt.Printf("defer：最后输出 %d \n", i)
	}
	fmt.Println("do something...")
}

func sample2() {
	//file.Open("file.txt")
	//defer file.Close()
}

// defer在匿名函数的应用
func sample3() {
	defer func() {
		fmt.Println("After defer 1")
	}()

	// 等价于
	//f := func() {
	//	fmt.Println("After defer 2")
	//}
	//defer f()

	fmt.Println("Before defer")
}