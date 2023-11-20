package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(sum(5, 5))
}

func sum(x int, y int) int {
	return x + y
}
