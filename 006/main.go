package main

import (
	"fmt"
	"math"
)

func main() {
	sumOfSquares := 0
	squaresOfSum := 0
	for i := 1; i <= 100; i++ {
		sumOfSquares += i * i
		squaresOfSum += i
	}
	squaresOfSum *= squaresOfSum

	fmt.Println(int(math.Abs(float64(sumOfSquares - squaresOfSum))))
}
