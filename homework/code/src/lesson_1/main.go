package main

import "fmt"

/*
*
给定一个字符串数组
[“I”,“am”,“stupid”,“and”,“weak”]
用 for 循环遍历该数组并修改为
[“I”,“am”,“smart”,“and”,“strong”]
*/
func main() {
	rawArray := []string{"I", "am", "stupid", "and", "weak"}
	targetArray := []string{"I", "am", "smart", "and", "strong"}
	var index int
	for index = 0; index < len(rawArray); index++ {
		rawArray[index] = targetArray[index]
	}
	fmt.Print(rawArray)
}
