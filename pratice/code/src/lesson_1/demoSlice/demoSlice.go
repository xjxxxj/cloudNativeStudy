package main

import "fmt"

func main() {
	myArray := []int{1, 2, 3, 4, 5}
	mySlice := myArray[1:3]
	fmt.Printf("mySlice: %+v\n", mySlice)
	fullSlice := myArray[:]
	remove3rdItem := deleteItem(fullSlice, 2)
	fmt.Printf("remove3rdItem: %+v\n", remove3rdItem)
}

func deleteItem(slice []int, i int) []int {
	// slice[i+1:]...是表示什么？从第i+1个元素开始截取到最后？
	return append(slice[:i], slice[i+1:]...)
}
