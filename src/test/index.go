package test

import "fmt"

func sum(a, b int) int {
	return a + b
}

func test() {
	fmt.Println(sum(1, 2))
}
