package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `default:"xqx"`
	Age  int    `json:"age"`
}

func (student Student) toString() {
	fmt.Printf("{\"name\":\"%s\", \"age\":\"%d\"}\n", student.Name, student.Age)
}

func printStudentInfo(student Student) {
	fmt.Printf("name :%s, age: %d\n", student.Name, student.Age)
}

func main() {
	student := Student{Name: "xjx", Age: 26}
	printStudentInfo(student)

	myStudent := reflect.TypeOf(student)
	name := myStudent.Field(0)
	tag := name.Tag.Get("default")
	fmt.Printf("default :%s\n", tag)

	student.toString()
}
