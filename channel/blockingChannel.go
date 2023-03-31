package main

import "fmt"

func main() {

	ch := make(chan string)

	go displayChannel(ch)

	fmt.Println("Not yet recieved value 1") // interesting output if you comment this line

	ch <- "Recieved Value"

	fmt.Println("Not yet recieved value 2") // interesting output if you comment this line
}

func displayChannel(ch chan string) {

	fmt.Println("Value recieved from Channel: ", <-ch)
}
