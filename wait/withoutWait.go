package main

import (
	"fmt"
	"time"
)

func runner1() {
	fmt.Println("\nI am first println")
}

func runner2() {
	fmt.Println("\nI am second println")
}

// This function will only execute when we have used time.Sleep
// so that the main function waits for goroutines to execute
func executeGoroutine() {
	go runner1()
	go runner2()
}

// This function executes
func executeAnyway() {
	runner1()
	runner2()
}

func main() {
	executeGoroutine()
	// executeGoroutine only executes because of the below line
	time.Sleep(time.Second)

	executeAnyway()
}
