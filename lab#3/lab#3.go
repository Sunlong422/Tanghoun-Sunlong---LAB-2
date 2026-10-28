package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter first integer: ")
	fmt.Scan(&a)

	fmt.Print("Etner second integer: ")
	fmt.Scan(&b)

	fmt.Println("")
	myAND(a, b)
	myOR(a, b)
	myXor(a, b)
	myNOT(a, b)
	myLeftShift(a, b)
	myRightShift(a, b)
}

func myAND(a, b int) {
	var myAnd = a & b
	fmt.Printf("%d & %d = %d\n", a, b, myAnd)
}

func myOR(a, b int) {
	var myOr = a | b
	fmt.Printf("%d | %d = %d\n", a, b, myOr)
}

func myXor(a, b int) {
	var myXor = a ^ b
	fmt.Printf("%d ^ %d = %d\n", a, b, myXor)
}

func myNOT(a, b int) {
	var myNot = (a != b)
	fmt.Printf("%d != %d = %t\n", a, b, myNot)

}

func myLeftShift(a, b int) {
	fmt.Printf("%d << 2 = %d\n", a, a<<2)
	fmt.Printf("%d << 2 = %d\n", b, b<<2)
}

func myRightShift(a, b int) {
	fmt.Printf("%d >> 2 = %d\n", a, a>>2)
	fmt.Printf("%d >> 2 = %d\n", b, b>>2)
}
