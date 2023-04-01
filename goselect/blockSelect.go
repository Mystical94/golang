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

	default:
		fmt.Println("Wait! Channels are getting ready")
	}

}

func channelNumber(number chan int) {

	time.Sleep(3 * time.Second)

	number <- 15
}

func channelMessage(message chan string) {

	time.Sleep(2 * time.Second)

	message <- "Code kar, procastination nahi"
}
