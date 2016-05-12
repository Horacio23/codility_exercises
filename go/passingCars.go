package main

import "fmt"

func main() {
	A := []int{0, 1, 0, 1, 1}

	fmt.Println("Amount of pairs:", Solution(A))
}

func Solution(A []int) int {
	count := 0
	mult := 0
	start := false
	for _, v := range A {
		if v == 0 {
			mult++
			start = true
		}

		if start && v == 1 {
			count = count + (1 * mult)

			if count > 1000000000 {
				return -1
			}

		}
	}

	return count
}
