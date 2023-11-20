package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	var randomX int = 10
	var randomY int = 5
	var sum int = sum(randomX, randomY)
	fmt.Println("Sum of", randomX, "and", randomY, "is", sum)
}

func sum(x int, y int) int {
	return x + y
}
