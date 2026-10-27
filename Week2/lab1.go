package main

import "fmt"

func lab1() {
	var a, b int

	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	fmt.Println("\n--- Assignment Operators ---")

	
	result := a
	fmt.Println("result = a  :", result)

	result += b
	fmt.Println("result += b :", result)

	result -= b
	fmt.Println("result -= b :", result)

	result *= b
	fmt.Println("result *= b :", result)

	if b != 0 {
		result /= b
		fmt.Println("result /= b :", result)

		result %= b
		fmt.Println("result %= b :", result)
	} else {
		fmt.Println("Cannot divide or modulo by zero")
	}
}
