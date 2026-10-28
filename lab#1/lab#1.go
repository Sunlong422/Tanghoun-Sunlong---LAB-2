package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter first integer: ")
	fmt.Scan(&a)

	fmt.Print("Enter second integer: ")
	fmt.Scan(&b)

	result := a
	fmt.Println("result = a:", result)

	result = a
	result += b
	fmt.Println("result = b:", result)

	result = a
	result -= b
	fmt.Println("result = b:", result)

	result = a
	result *= b
	fmt.Println("result = b:", result)

	if b != 0 {
		result = a
		result /= b
		fmt.Println("result = b:", result)

		result = a
		result %= b
		fmt.Println("result %%= b:", result)
	} else {
		fmt.Println("\n%= and /= operator: skipped (division by zero)")
	}
}
