package main

import (
	"fmt"
	"strconv"
)


func main() {
	max := 1000
	maxPalindrome := 0
	for i := 0; i < max; i++ {
		for j := 0; j < max; j++ {
			result := i * j
			if result < maxPalindrome {
				continue
			}

			resultStr := strconv.Itoa(result)
			reverseStr := ""
			for k := len(resultStr) - 1; k >= 0; k-- {
				reverseStr += string(resultStr[k])
			}
			if resultStr == reverseStr {
				if result > maxPalindrome {
					maxPalindrome = result
				}
			}
		}
	}
	fmt.Println(maxPalindrome)
}
