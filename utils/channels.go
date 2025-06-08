package main

import (
	"fmt"
)

func calculateSquares(number int, squareStream chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareStream <- sum
}

func calculateCubes(number int, cubesStream chan int) {
	sum := 0

	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubesStream <- sum
}

func main() {
	number := 12345
	squaresChannel := make(chan int)
	cubesChannel := make(chan int)

	go calculateCubes(number, cubesChannel)
	go calculateSquares(number, squaresChannel)

	squares, cubes := <-squaresChannel, <-cubesChannel
	fmt.Println("Squares:", squares)
	fmt.Println("Cubes:", cubes)
	fmt.Println("Sum:", squares+cubes)
}