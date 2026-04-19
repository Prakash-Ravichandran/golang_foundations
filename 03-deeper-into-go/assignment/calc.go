package main

import "fmt"

type intArr []int

func main() {
	nums := intArr{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range nums {
		if num%2 == 0 {
			fmt.Printf("num is even %v\n", num)
		} else {
			fmt.Printf("num is odd %v\n", num)
		}
	}

}