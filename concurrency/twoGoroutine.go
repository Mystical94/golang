package main

import (
	"fmt"
	"time"
)

func display(message string) {
	for i := 1; i < 10; i++ {
		fmt.Println(message)

		// to sleep the process for 1 sec
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go display("process 1")

	display("Process 2")
}
