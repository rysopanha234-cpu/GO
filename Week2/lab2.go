package main

import "fmt"

func lab2() {
	var a,b int 

	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter Second number: ")
	fmt.Scan(&b)

fmt.Println("\n--- Logical operators ---")

	bothPositive := (a>0) && ( b>0)
	oneGreater := (a>b)|| (a<b)
	notEqual := !(a==b)

	fmt.Println("Both numbers are positive:", bothPositive)
    fmt.Println("One number is greater than the other:", oneGreater)
    fmt.Println("The numbers are not equal:", notEqual)


}