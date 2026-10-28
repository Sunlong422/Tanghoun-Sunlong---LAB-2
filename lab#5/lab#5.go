package main

import (
	"fmt"
	"strconv"
)

func main() {
	var text string

	fmt.Print("Enter your text: ")
	fmt.Scan(&text)

	bytes := []byte(text)
	fmt.Print("Binary: ")
	binaryDis(bytes)
	fmt.Print("\n")
	fmt.Print("Hexadecimal: ")
	hexaDis(bytes)
	fmt.Print("\n")
	fmt.Print("Base64: ")
	base64Dis(bytes)
}

// display as binary
func binaryDis(s []byte) {

	for i := 0; i < len(s); i++ {
		n := s[i]
		binSlice := []int{}
		for n > 0 {
			remainder := n % 2
			binSlice = append(binSlice, int(remainder))
			n = n / 2
		}
		for i := len(binSlice) - 1; i >= 0; i-- {
			fmt.Print(binSlice[i])
		}
		fmt.Print("\t")
	}
}

// display as hexadecimal
func hexaDis(s []byte) {

	for i := 0; i < len(s); i++ {
		n := s[i]
		binSlice := []string{}
		for n > 0 {
			remainder := n % 16

			if remainder < 10 {
				binSlice = append(binSlice, strconv.Itoa(int(remainder)))
			} else if remainder == 10 {
				binSlice = append(binSlice, "A")
			} else if remainder == 11 {
				binSlice = append(binSlice, "B")
			} else if remainder == 12 {
				binSlice = append(binSlice, "C")
			} else if remainder == 13 {
				binSlice = append(binSlice, "D")
			} else if remainder == 14 {
				binSlice = append(binSlice, "E")
			} else if remainder == 15 {
				binSlice = append(binSlice, "F")
			}

			n = n / 16
		}
		for i := len(binSlice) - 1; i >= 0; i-- {
			fmt.Print(binSlice[i])
		}
		fmt.Print("\t")
	}
}

// display as base64
func base64Dis(s []byte) {
	base64Chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	for i := 0; i < len(s); i++ {
		n := s[i]
		binSlice := []string{}

		if n == 0 {
			binSlice = append(binSlice, "A")
		}

		for n > 0 {
			remainder := n % 64
			binSlice = append(binSlice, string(base64Chars[remainder]))
			n = n / 64
		}

		for j := len(binSlice) - 1; j >= 0; j-- {
			fmt.Print(binSlice[j])
		}
		fmt.Print("\t")
	}
}
