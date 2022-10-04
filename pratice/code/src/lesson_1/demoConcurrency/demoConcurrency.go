package main

import (
	"fmt"
	"strconv"
	"time"
)

func prod(ch chan<- string, productNum int) {
	for num := 0; num < productNum; num++ {
		//n := rand.Intn(3)
		//time.Sleep(time.Duration(n))
		fmt.Printf("product:%d\n", num)
		ch <- "prod num:" + strconv.Itoa(num)
	}
	close(ch)
}

func consume(ch <-chan string) {
	for {
		num, notClosed := <-ch
		if !notClosed {
			break
		}
		fmt.Printf("consume:%s\n", num)
	}
}

func main() {
	for i := 0; i < 1000; i++ {
		go fmt.Println("num:", i)
	}
	time.Sleep(time.Second)
	fmt.Println("=====================================================================")
	var ch = make(chan string)
	go prod(ch, 10)
	go consume(ch)
	time.Sleep(10)
}
