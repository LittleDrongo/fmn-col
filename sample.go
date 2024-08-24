package main

import (
	"fmt"
	"queue/collections"
)

func sampleQueueString() {

	stringQueue := collections.NewQueue[string]()
	fmt.Println(stringQueue.Get())

}
