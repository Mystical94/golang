package main

import "fmt"

func test1() {

	var a int = 10

	if a%2 == 0 {
		fmt.Printf("%d is an even number", a)
	} else {
		fmt.Printf("%d is an odd number", a)
	}
}

func test2() {

	fmt.Print("\nEnter Number:")
	var input int
	fmt.Scanln(&input)
	fmt.Printf("\nEntered Number: %d", input)

	if input%2 == 0 {
		fmt.Printf("\n%d is an even number", input)
	} else {
		fmt.Printf("\n%d is an odd number", input)
	}
}

func main() {
	test1()
	test2()
}
