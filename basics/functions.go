package basics

import "fmt"

func main() {
	x, y := 7, 5
	var result int

	result = applyOperation(x, y, add)
	fmt.Println(result)

	result = applyOperation(x, y, sabtract)
	fmt.Println(result)

	result = applyOperation(x, y, multiply)
	fmt.Println(result)

	result = createMultiplier(5)(3)
	fmt.Println(result)

	quotient, remainder := divideInts(x, y)
	fmt.Printf("\tQuotient: %d,\n \tRemainder: %d\n\n", quotient, remainder)
}

func applyOperation(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		fmt.Printf("x is: %d, factor is: %d\n", x, factor)
		return x * factor
	}
}

func add(x, y int) int {
	return x + y
}

func sabtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divideInts(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
