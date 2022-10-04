package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
• 基于 Channel 编写一个简单的单线程生产者消费者模型
• 队列：
队列长度10，队列元素类型为 int
• 生产者：
每1秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
• 消费者：
每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
*/

func prod(ch chan<- int) {
	for {
		time.Sleep(time.Second)
		item := rand.Intn(10)
		ch <- item
		fmt.Printf("product: %d\n", item)
	}
}

func consume(ch <-chan int) {
	for {
		time.Sleep(time.Second)
		item, notClosed := <-ch
		if !notClosed {
			break
		}
		fmt.Printf("consume: %d\n", item)
	}
}

func main() {
	/*	// 尝试channel是空时，消费者获取到的是什么: 获取时直接阻塞
		ch1 := make(chan int, 10)
		go consume(ch1)
		time.Sleep(10 * time.Second)

		// 尝试生产者放满channel时会怎么样: 放入时会阻塞
		ch2 := make(chan int, 2)
		go prod(ch2)
		time.Sleep(15 * time.Second)*/

	ch := make(chan int, 10)
	go prod(ch)
	go consume(ch)
	time.Sleep(30 * time.Second)

}
