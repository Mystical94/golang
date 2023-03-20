package main

import "fmt"

func main() {
	fmt.Printf("\nEnter Number: ")
	var input int
	fmt.Scanln(&input)

	switch input {
	case 1:
		fmt.Println("The case chosen was 1")
		//fmt.Println("The case chosen was 1"); fallthrough;
		//fallthrough is used at end of each case statement to print all the cases present below
	case 2:
		fmt.Println("The case chosen was 2")
	case 3:
		fmt.Println("The case chosen was 3")
	case 4:
		fmt.Println("The case chosen was 4")
	case 5:
		fmt.Println("The case chosen was 5")
	default:
		fmt.Println("The case chosen was not present")
	}

}
