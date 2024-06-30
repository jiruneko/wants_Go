package main

import "fmt"

const x int64 = 10

const (
	idKey   = "id"
	namekey = "name"
)

const z = 20 * 20

func main() {
	const y = "Hello"

	fmt.Println(x)
	fmt.Println(y)

	x := x + 1
	fmt.Println(x)
	fmt.Println(y)
}
