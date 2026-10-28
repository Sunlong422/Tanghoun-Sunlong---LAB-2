package main

import "fmt"

func main() {
	var text string
	var key int

	fmt.Print("Enter your text: ")
	fmt.Scanln(&text)

	fmt.Print("Enter key (0-255): ")
	fmt.Scan(&key)

	// Encrypt
	ciphertext := xorEncrypt(text, byte(key))
	fmt.Println("\nEncrypted (raw bytes):", []byte(ciphertext))

	// Decrypt — using the SAME function with the SAME key
	decrypted := xorEncrypt(ciphertext, byte(key))
	fmt.Println("Decrypted:", decrypted)
}

func xorEncrypt(text string, key byte) string {
	result := make([]byte, len(text))

	for i := 0; i < len(text); i++ {
		result[i] = text[i] ^ key
	}

	return string(result)
}
