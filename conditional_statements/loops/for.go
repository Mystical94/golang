package main

import "fmt"

func test1() {
	for a := 0; a < 11; a++ {
		fmt.Print(a, "\n")
	}
}

func condition() {
	sum := 1
	for sum < 100 {
		sum += sum
		fmt.Println(sum)
	}
}

func main() {
	condition()
}
