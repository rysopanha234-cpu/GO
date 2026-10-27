package main

import "fmt"

func add(a, b int) {
	fmt.Println("Result:", a+b)
}

func sub(a, b int) {
	fmt.Println("Result:", a-b)
}

func mul(a, b int) {
	fmt.Println("Result:", a*b)
}

func div(a, b int) {
	if b == 0 {
		fmt.Println("error")
	} else {
		fmt.Println("Result:", a/b)
	}
}

func mod(a, b int) {
	if b == 0 {
		fmt.Println("error")
	} else {
		fmt.Println("Result:", a%b)
	}
}

func lab4() {
	var choice int
	var a, b int

	for {
		fmt.Println("\n Mini Calculator")
		fmt.Println("1) Add 2) Sub 3) Mul 4) Div 5) Mod 6) Exit")
		fmt.Print("Choose: ")
		fmt.Scan(&choice)

		if choice == 6 {
			fmt.Println("Exiting")
			break
		}

		if choice < 1 || choice > 6 {
			fmt.Println("Error: Invalid choice")
			continue
		}

		fmt.Print("Enter a: ")
		fmt.Scan(&a)
		fmt.Print("Enter b: ")
		fmt.Scan(&b)

		switch choice {
		case 1:
			add(a, b)
		case 2:
			sub(a, b)
		case 3:
			mul(a, b)
		case 4:
			div(a, b)
		case 5:
			mod(a, b)
		}
	}
}
