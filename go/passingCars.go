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
			// every time a zero is found, we count each 1 as many times as we have encountered zeros
			mult++
			start = true
		}

		if start && v == 1 {
			// add the current count to the multiplication of all the zeroes we have found
			count = count + (1 * mult)

			if count > 1000000000 {
				return -1
			}

		}
	}

	return count
}
