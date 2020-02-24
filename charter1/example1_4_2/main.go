package main

import "fmt"

func main() {
	//
	student := new(Student)
	student.name = "张三"
	student.age = 18
	student.speciality = "篮球"
	name, age := student.getNameAndAge()
	speciality := student.getSpeciality()
	fmt.Println(name, age, speciality)
}

type Person struct {
	name string
	age int
}

func (person Person) getNameAndAge() (string, int) {
	return person.name, person.age
}

type Student struct {
	Person
	speciality string
}

func (student Student) getSpeciality() string {
	return student.speciality
}

