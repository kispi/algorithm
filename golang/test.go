package main

import (
	"fmt"
)

func main() {
	test()
}

func test() {
	jakadIndex := Jakad("pneumonoconiosis", "pneumonoultramicroscopicsilicovolcanoconiosis", true)
	fmt.Println("Jakad index: ", jakadIndex)
}
