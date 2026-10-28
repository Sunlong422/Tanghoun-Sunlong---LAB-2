package main

import "fmt"

func main() {
	loop := true
	var input int
	var a, b int
	for loop {
		fmt.Println("===== MINI CALCULATOR =====")
		fmt.Println("1) Add 2) Sub 3) Mul 4) Div 5) Mod 6) Exit")
		fmt.Print("Choose: ")
		fmt.Scan(&input)

		if input != 6 {
			fmt.Print("Enter a: ")
			fmt.Scan(&a)

			fmt.Print("Enter b: ")
			fmt.Scan(&b)

			switch input {
			case 1:
				fmt.Println("Result:", addNum(a, b))
			case 2:
				fmt.Println("Result:", subNum(a, b))
			case 3:
				fmt.Println("Result:", mulNum(a, b))
			case 4:
				fmt.Println("Result:", divNum(a, b))
			case 5:
				fmt.Println("Result:", modNum(a, b))
			}
		} else {
			loop = false
		}
	}
}

func addNum(a, b int) int {
	return a + b
}

func subNum(a, b int) int {
	return a - b
}

func mulNum(a, b int) int {
	return a * b
}

func divNum(a, b int) int {
	return a / b
}

func modNum(a, b int) int {
	return a % b
}
