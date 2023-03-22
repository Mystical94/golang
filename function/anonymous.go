package main

import "fmt"

func test1() {
	// anonymous function
	var greet = func() {
		fmt.Println("Anonymous Function")
	}

	greet()
	greet()
	greet()
	greet()
	greet()
}

func test2() {
	// anonymous function with arguments
	var sum = func(n1, n2 int) {
		sum := n1 + n2
		fmt.Println("\nSum: ", sum)
	}

	sum(5, 3)
	sum(10, 15)
	sum(4, 9)
}

func test3() {
	// anonymous function with arguments
	var sum = func(n1, n2 int) int {
		sum := n1 + n2
		return sum
	}

	result := sum(5, 3)
	fmt.Println("\nSum: ", result)
}

// regular function to calculate square of numbers
func findSquare(num int) int {
	return num * num
}

func test4() {
	// anonymous function that returns sum of numbers
	sum := func(n1, n2 int) int {
		return n1 + n2
	}
	// function call
	result := findSquare(sum(5, 4))
	fmt.Println(result)
}

func displayNumber() func(number int) int {
	//number := 10
	return func(number int) int {
		number++
		return number
	}
}

func test5() {
	a := displayNumber()
	fmt.Println(a(10))
}

func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	test5()
}
