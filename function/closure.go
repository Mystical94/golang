package main

import "fmt"

func calculate() func() int {
	num := 1

	//returns inner function
	return func() int {
		num = num + 2
		return num
	}
}

func main() {

	odd := calculate()

	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(odd())

	odd2 := calculate()

	fmt.Println(odd2())
}
