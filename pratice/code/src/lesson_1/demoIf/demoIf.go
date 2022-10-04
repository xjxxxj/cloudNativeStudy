package main

import "fmt"

func main() {
	var name = "xjx"
	if name == "xxj" {
		fmt.Println("girl, 25 years old.")
	} else if name == "xqx" {
		fmt.Println("man, 26 years old.")
	} else {
		fmt.Println("love, 10 years old.")
	}

	var age = 18
	if age := age + 10; age < 30 {
		fmt.Println("less than 30 years old.")
	}
}
