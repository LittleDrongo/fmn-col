package main

import (
	"fmt"
	"queue/collections"
)

func main() {

	sq := collections.NewQueue[int]()
	sq.Put(0)
	sq.Put(5)
	// stringQueue.Put("sss")

	fmt.Println(sq.Get())

	fmt.Println(sq.Get())

	fmt.Println(sq.Get())

}
