package main

import "fmt"

func main() {
	message := ""
	message = "Hello world!"
	fmt.Println(message)

	var testVar string = "1"
	fmt.Println(testVar)

	const testConst int = 2
	fmt.Println(testConst)
}

// go run main.go
// go build main.go
// ./main
