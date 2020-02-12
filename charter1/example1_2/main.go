package main

import "fmt"

func main() {
	doSwap()
	doAdd1()
	doAdd2()
	doSum()
	anonymous()
}

func swap(a int, b int) (int, int)  {
	return b,a
}

func doSwap() {
	a := 1
	b := 3
	a, b = swap(a, b)
	fmt.Println("------- doAdd ------")
	fmt.Printf("doSwap：%d %d \n\n", a, b)
}

// 值传递
func add1(a int) int {
	a += 1
	return a
}

func doAdd1() {
	x := 1
	y := add1(x)
	fmt.Println("------- doAdd1 ------")
	fmt.Printf("y：%d \n", y) // y 输出值为2
	fmt.Printf("x: %d \n\n", x) // x 的输出值为1，并没有受函数add1()操作影响
}


//指针传递
//1. 传指针使得多个函数能操作同一个对象。
//2. 传指针比较轻量级 (8bytes),只是传内存地址，我们可以用指针传递体积大的结构体。
//	 在每次copy上面就会花费相对较多的系统开销（内存和时间）。所以当你要传递大的结构体的时候，用指针是一个明智的选择。
//3. Go语言中channel，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。（注：若函数需改变slice的长度，则仍需要取地址传递指针）
func add2(a *int) int {
	*a = *a + 1
	return *a
}

// 指针传递，将x指针（引用）传递至 add2(&x)，变量经过了 add2() 内的逻辑计算后，最终被修改
func doAdd2() {
	x := 1
	y := add2(&x)
	fmt.Println("------- doAdd2 ------")
	fmt.Printf("y：%d \n", y) // 此时 y 输出2
	fmt.Printf("x: %d \n\n", x) // x的值也变为了 2，达到了修改变量 x 值的目的
}

// slice
func sum(num []int) int {
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	return sum
}

func doSum() {
	num := []int{1, 2, 3, 4, 5}
	fmt.Println("------- doSum ------")
	fmt.Println(sum(num))
}

// 匿名函数
func anonymous() {
	f := func(a int, b int) int {
		return a + b
	}
	fmt.Println("------- anonymous ------")
	fmt.Println(f(2,3))
}