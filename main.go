package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello, World!")

	var randomX int = (rand.Int() * 100)
	var randomY int = (rand.Int() * 100)
	var sum int = sum(randomX, randomY)
	fmt.Println("Sum of", randomX, "and", randomY, "is", sum)
}

func sum(x int, y int) int {
	return x + y
}
