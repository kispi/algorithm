package main

import (
	"fmt"

	"./golang"
)

func main() {
	test()
}

func test() {
	jakadIndex := golang.Jakad("pneumonoconiosis", "pneumonoultramicroscopicsilicovolcanoconiosis", true)
	fmt.Println("Jakad index: ", jakadIndex)
}
