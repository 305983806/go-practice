package main

import "fmt"

func main() {
	sample1()
	sample2()
	sample3()
	sample4()
}

// if
func sample1()  {
	x := 100
	if x > 100 {
		fmt.Println("满分，你太棒了！")
	} else if x >= 90 {
		fmt.Println("90分以上，下次争取100分。")
	} else {
		fmt.Println("抱歉，低于90分，不及格。")
	}
}

// for
func sample2() {
	x := 0
	sum := 0
	for x <= 100 {
		sum += x
		x++
	}
	fmt.Println(sum)
}

/*
switch
注意：
	1. 无需加 break;
	2. 如想执行完一个case后，继续向下执行，可以添加 fallthrough
 */
func sample3()  {
	x := 2
	switch x {
		case 1:
			println("张三")
			//fallthrough
		case 2:
			println("李四")
			//fallthrough
		case 3:
			println("王五")
			//fallthrough
		default:
			println("I don't know")
	}
}

/*
range：用于遍历容器中的元素（相当于java 或 php 中的 foreach）
其中 i 代表 key，v 代表 value。
如果只需要遍历 key 或 value 中的一项，只需要使用 "_" 占位符，忽略掉这个变量即可
 */
func sample4()  {
	//x := [5]int{1, 2, 3, 4, 5}

	x := make(map[string]int)
	x["张三"] = 7
	x["李四"] = 9
	x["王五"] = 3

	for i, _ := range x {
		//fmt.Println(i, v)
		fmt.Println(i)
	}
}