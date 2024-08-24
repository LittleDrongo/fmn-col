package main

import (
	"fmt"
	"queue/collections"
)

func main() {

	queue := collections.NewQueue()

	fmt.Println("queue size is: ", queue.Size())

	queue.Put(1)

	queue.Put(2)

	queue.Put("world")

	queue.Put('x')

	fmt.Println("queue size is: ", queue.Size())

	queue.Put(func() {
		fmt.Printf("func test")

	})

	for item := queue.Get(); item != nil; item = queue.Get() {

		switch v := item.(type) {
		case func():
			v()
		default:
			fmt.Println(v)
		}
	}
}
