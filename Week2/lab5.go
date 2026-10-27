package main

import (
	"encoding/base64"
	"fmt"
)

func lab5() {
	var text string

	fmt.Print("Enter text: ")
	fmt.Scan(&text)

	fmt.Print("Binary: ")
	for _, c := range text {
		fmt.Printf("%08b ", c)
	}
	fmt.Println()

	fmt.Printf("Hexadecimal: %x\n", text)
	
	b64 := base64.StdEncoding.EncodeToString([]byte(text))
	fmt.Println("Base64:", b64)
}
