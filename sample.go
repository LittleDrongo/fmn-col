package main

import (
	"queue/collections"
)

func main() {

	q := collections.NewQueue[any]()

	q.Put(30)
	q.Put("dd")
	q.Put('d')
	q.Put("dd")

	// // element := q.Get()
	// fmt.Println(
	// 	q.Get(),
	// 	q.Get(),
	// 	q.Get(),
	// 	q.Get(),
	// 	q.Get(),
	// )

}
