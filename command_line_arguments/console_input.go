/*
How to run this file

go build console_input.go
After running above build command
run ./console_input [your values]
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, arg string
	for i := 1; i < len(os.Args); i++ {
		s += arg + os.Args[i] + " "
	}

	fmt.Println(s)
}
