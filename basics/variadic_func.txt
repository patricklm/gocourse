package main

import "fmt"

func main() {
	var result int
	values := []int{4, 3, 6, 2, 9}

	result = sum(values...)
	fmt.Println("The sum of all numbers:", result)

	result = sum(5,7)
	fmt.Println("The sum of all numbers:", result)
}

func sum(nums ...int) int {
	total := 0
	for _,n := range nums {
		total += n
	}
	return total
}
