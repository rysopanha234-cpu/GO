package main

import "fmt"

func myAND(a, b int) {
	fmt.Println("AND  =", a&b)
}

func myOR(a, b int) {
	fmt.Println("OR  =", a|b)
}

func myXor(a, b int) {
	fmt.Println("XOR  =", a^b)
}

func myNOT(a, b int) {
	fmt.Println("NOT a =", ^a, ", NOT b =", ^b)
}

func myShift(a, b int) {
	fmt.Println("Left shift a  =", a<<1)
	fmt.Println("Right shift a  =", a>>1)
	fmt.Println("Left shift b  =", b<<1)
	fmt.Println("Right shift b =", b>>1)
}

func myAssignment(a, b int) {
	x := a
	fmt.Println("x = a  :", x)

	x += b
	fmt.Println("x += b :", x)

	x -= b
	fmt.Println("x -= b :", x)

	x *= b
	fmt.Println("x *= b :", x)

	if b != 0 {
		x /= b
		fmt.Println("x /= b :", x)

		x %= b
		fmt.Println("x %= b :", x)
	} else {
		fmt.Println("Cannot divide by zero")
	}
}

func lab3() {
	var a, b int

	fmt.Print("Enter first number: ")
	fmt.Scan(&a)

	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	fmt.Println("\n--- Bitwise Operations ---")
	myAND(a, b)
	myOR(a, b)
	myXor(a, b)
	myNOT(a, b)
	myShift(a, b)

	fmt.Println("\n--- Assignment Operations ---")
	myAssignment(a, b)
}
