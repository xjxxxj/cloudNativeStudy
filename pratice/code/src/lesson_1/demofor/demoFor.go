package main

import "fmt"

func main() {
	var sum = 0
	var maxNum = 100
	for i := 0; i < maxNum; i++ {
		sum += i
	}
	fmt.Printf("0 + 1 + ... + %d = %d\n", maxNum, sum)

	var k = 0
	sum = 0
	for k < maxNum {
		sum += k
		k++
	}
	fmt.Printf("0 + 1 + ... + %d = %d\n", maxNum, sum)

	// for string
	var target = "love"
	for _, char := range target {
		fmt.Println(char)
	}

	// for array
	var names = []string{"xqx", "xxj"}
	for index, name := range names {
		fmt.Printf("no:%d,name:%s\n", index, name)
	}
	// for map
	myInfo := map[string]string{
		"xxj": "25",
		"xqx": "26",
	}
	for name, age := range myInfo {
		fmt.Printf("name: %s, age ：%s\n", name, age)
	}
}
