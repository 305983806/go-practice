package main

import "fmt"

func main() {
	// 接口定义与实现
	// 一个类只要实现了接口的所有函数，就说明这个类实现了这个接口
	var bird Animal = new(Bird)
	bird.run()
	bird.fly()

	// 将一个接口赋值给另一个接口
	// 包含函数多的接口对象，可以赋值给包含函数较少的接口对象
	var bird2 Animal2 = bird
	bird2.fly()

	// 类型查询
	var v1 interface{}
	v1 = "张三"
	if v, ok := v1.(string); ok {
		fmt.Println(v, ok)
	}

	// v1.(type)，只能在 switch中使用
	switch v1.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("float32")
	case float64:
		fmt.Println("float64")
	}
}

type Animal interface {
	run()
	fly()
}

type Animal2 interface {
	fly()
}

type Bird struct {

}

func (bird Bird) run()  {
	fmt.Println("Bird is running!")
}

func (bird Bird) fly()  {
	fmt.Println("Bird is flying!")
}