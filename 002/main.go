package main

import "fmt"

const max = 4000000

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n - 1) + fibonacci(n - 2)
}

func main() {
	sum := 0
	for n, f := 0, 0; f < max; n++ {
		f = fibonacci(n)
		if f % 2 == 0 {
			sum += f
		}
	}
	fmt.Println(sum)
}
