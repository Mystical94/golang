package main

//import the custom package calculator
import (
	"fmt"
	//"package/calculator"
	"custom_package/package/calculator"
)

func main() {

	num1 := 9
	num2 := 5

	//use the add function of calculator package
	fmt.Println(calculator.Add(num1, num2))

	//use the add function of calculator package
	fmt.Println(calculator.Substract(num1, num2))
}
