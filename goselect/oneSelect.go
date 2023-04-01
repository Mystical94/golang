package main

import (
	"fmt"
	"time"
)

func main() {

	number := make(chan int)
	message := make(chan string)

	go channelNumber(number)
	go channelMessage(message)

	select {
	case firstChannel := <-number:
		fmt.Println("Channel Data:", firstChannel)

	case secondChannel := <-message:
		fmt.Println("Channel Data:", secondChannel)
	}
}

func channelNumber(number chan int) {
	number <- 15
}

func channelMessage(message chan string) {

	// sleep the process by 2 seconds
	time.Sleep(2 * time.Second)

	message <- "Learning Go Select"
}
