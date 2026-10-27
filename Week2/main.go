package main

import "fmt"

func main() {
	var choice int

	for {
		fmt.Println("===== Choose a lab to run (1-6) =====")
		fmt.Println("1. Assignment operators")
		fmt.Println("2. Logical operators")
		fmt.Println("3. Bitwise operators")
		fmt.Println("4. Calculator")
		fmt.Println("5. Binary, Hex, Base64 encoding")
		fmt.Println("6. XOR encryption/decryption")
		fmt.Println("7. Exit")
		fmt.Print("Choose a lab: ")

		fmt.Scan(&choice)

		if choice == 7 {
			fmt.Println("\n(˶ᵔᗜᵔ˶)ﾉﾞ!")
			break
		}

		switch choice {
		case 1:
			lab1()
		case 2:
			lab2()
		case 3:
			lab3()
		case 4:
			lab4()
		case 5:
			lab5()
		case 6:
			lab6()
		default:
			fmt.Println("Invalid choice")
		}

		fmt.Println()
	}
}

