package main

import "fmt"

func main() {
	// panic
	sample1()

	// type
	sample2()

	// recover
	firstName := "Chen"
	//lastName := "Peidong"
	sample3(&firstName, nil)
}

func sample1() {
	fmt.Println("------- sample1 ------")
	defer func() {
		fmt.Println("After defer")
	}()

	//panic("An error has occurred!")

	fmt.Println("Before defer")
}

func sample2() {
	fmt.Println("------- sample2 ------")
	type sum func(a, b int) int
	var getSum sum = func(a, b int) int {
		return a + b
	}
	fmt.Println(getSum(2, 3))
}

// recover() 是一个内建函数，用于重新获得 panic 协程的控制，可以取到 panic 的错误信息，并且停止 panic 续发事件（Panicking Sequence），
// 程序运行恢复正常。
// PS: 只有在 defer 函数的内部使用 recover()，如果在延迟函数的外部调用 recover，就不能停止 panic 续发事件。
func sample3(firstName *string, lastName *string) {
	fmt.Println("------- sample3 ------")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover from ", r)
		}
	}()

	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("return normally from main")
}