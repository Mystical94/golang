package main

import (
	"fmt"
)

func main() {

	// create channel of integer type
	number := make(chan int)

	// access type and value of channel
	fmt.Printf("Channel Type: %T\n", number)
	fmt.Printf("Channel Value: %v", number)
}
