package main

import "fmt"

func main() {

	//create two channels
	number := make(chan int)
	message := make(chan string)

	// function call with goroutine
	go channelNumber(number)
	go channelMessage(message)

	// selects and execute a channel
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
	message <- "Learning Go Select"
}
