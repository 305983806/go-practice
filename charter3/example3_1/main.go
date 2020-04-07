package main

import (
	"encoding/json"
	"fmt"

	"github.com/pquerna/ffjson/ffjson"
)

// Student 学生类型
type Student struct {
	// 姓名
	Name string `json:"student_name"`
	// 年龄
	Age int
}

func main() {
	var (
		arr     [5]int
		m       map[string]float64
		student Student
		s4      interface{}
	)

	// JSON 编码
	// func Marshal(v interface{}) ([]type, error)

	// JSON 解码
	// func Unmarshal(data []byte, v interface{}) error

	arr2Json(arr)
	map2Json(m)
	obj2Json(student)
	json2Arr(student, s4)
}

func arr2Json(arr [5]int) {
	arr = [5]int{1, 2, 3, 4, 5}
	s, err := json.Marshal(arr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))
}

func map2Json(m map[string]float64) {
	m = make(map[string]float64)
	m["zhangsan"] = 100
	s, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))
}

func obj2Json(student Student) {
	student = Student{"zhangsan", 16}
	s, err := json.Marshal(student)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))

	// 如果希望对象的 JSON 可以顺利打印，struct 的字段名必须大写
}

func json2Arr(student Student, s4 interface{}) {
	student = Student{"zhangsan", 16}
	s, err := json.Marshal(student)
	if err != nil {
		panic(err)
	}

	ffjson.Unmarshal([]byte(s), &s4)
	if err != nil {
		print(err)
	}
	fmt.Printf("%v\n", s4)

}
