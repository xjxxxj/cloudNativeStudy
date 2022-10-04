package main

import "fmt"

func main() {
	var name = "xjx"
	switch name {
	case "xxj":
		fmt.Println("girl, 25 years old.")
		break
	case "xqx":
		fmt.Println("man, 26 years old.")
		break
	default:
		fmt.Println("love, 10 years old.")
	}
}
