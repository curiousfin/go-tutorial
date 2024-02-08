package main

import (
	"fmt"
	"reflect"
)

func main() {
	message := ""
	message = "Hello world!"
	fmt.Println(message)
	fmt.Println(reflect.TypeOf(message))

	var testVar string = "1"
	fmt.Println(testVar)

	const testConst int = 2
	fmt.Println(testConst)
}

// go run main.go
// go build main.go
// ./main
