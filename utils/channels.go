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
	close(squareStream)
}

func calculateCubes(number int, cubesStream chan int) {
	sum := 0

	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubesStream <- sum
	close(cubesStream)
}

func main() {
	number := 589
	squaresChannel := make(chan int)
	cubesChannel := make(chan int)

	go calculateCubes(number, cubesChannel)
	go calculateSquares(number, squaresChannel)

	for {
		sq, ok := <-squaresChannel
		if !ok {
			break
		}
		cu, ok := <-cubesChannel
		if !ok {
			break
		}
		fmt.Println("Squares:", sq)
		fmt.Println("Cubes:", cu)
		fmt.Println("Sum:", sq+cu)
	}

	
}