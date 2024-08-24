package main

import (
	"fmt"
	"queue/collections"
)

func main() {

	sq := collections.NewQueue[string]()
	sq.Put("ds")
	sq.Put("dddd")
	// stringQueue.Put("sss")

	fmt.Println(sq.Get())

	fmt.Println(sq.Get())

	fmt.Println(sq.Get())

}
