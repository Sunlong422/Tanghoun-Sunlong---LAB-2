package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter first integer: ")
	fmt.Scan(&a)

	fmt.Print("Enter Second integer: ")
	fmt.Scan(&b)

	bothPositive := (a > 0) && (b > 0)
	greaterThan := (a > b) || (a < b)
	notE := a != b

	fmt.Println("Both positive (&&): ", bothPositive)
	fmt.Println("one greater than the other(||): ", greaterThan)
	fmt.Println("not equal (!): ", notE)
}
