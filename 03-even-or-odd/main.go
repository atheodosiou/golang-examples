package main

import "fmt"

func main() {
	numbers := []int{}

	for i := 1; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	for i := range numbers {
		if numbers[i]%2 != 0 {
			fmt.Printf("Number %d is odd\n", numbers[i])
		} else {
			fmt.Printf("Number %d is even\n", numbers[i])
		}
	}
}
