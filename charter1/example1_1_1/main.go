package main

import "fmt"

func main() {
	example1()
	example2()
	example3()
}

type Person struct {
	name string
	age int
}

type Student struct {
	Person
	num string
}
type StudentX struct {
	Person
	age int
	num string
}

func example1() {
	fmt.Println("------- sample1 ------")
	// 以下两种写法效果是一样的
	//person := Person{"cp.Chen", 18}
	person := Person{name: "cp.Chen", age: 18}
	fmt.Printf("姓名: %s，年龄：%d\n", person.name, person.age)
}

func example2()  {
	fmt.Println("------- sample2 ------")
	student := Student{
		Person{"cp.Chen", 18},
		"30",
	}
	fmt.Printf("姓名: %s，年龄：%d，学号：%s \n", student.name, student.age)
}

func example3() {
	fmt.Println("------- sample3 ------")
	studentX := StudentX{Person{"cp.Chen", 18}, 6, "30"}
	fmt.Printf("Person.age：%d，StudentX.age：%d\n", studentX.Person.age, studentX.age)
	fmt.Printf("%v\n", studentX)
}