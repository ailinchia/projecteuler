package main

import "fmt"

func main() {
	for n := 20; ; n += 2 {
		divisible := true
		for i := 2; i <= 20; i++ {
			if n % i != 0 {
				divisible = false
				break
			}
		}
		if divisible {
			fmt.Println(n)
			break
		}
	}
}
