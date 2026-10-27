package main

import "fmt"

func xorEncrypt(text string, key byte) string {
	out := []byte(text)
	for i := 0; i < len(out); i++ {
		out[i] = out[i] ^ key
	}
	return string(out)
}

func lab6() {
	var message, keyInput string

	fmt.Print("Enter plaintext message: ")
	fmt.Scan(&message)

	fmt.Print("Enter key (single character): ")
	fmt.Scan(&keyInput)

	key := keyInput[0]

	// 1. Encrypt
	ciphertext := xorEncrypt(message, key)
	fmt.Printf("Ciphertext (Hex) : %x\n", ciphertext)
	fmt.Println("Ciphertext (Raw) :", ciphertext)

	// 2. Decrypt using the same function
	decrypted := xorEncrypt(ciphertext, key)
	fmt.Println("Decrypted message:", decrypted)
}
