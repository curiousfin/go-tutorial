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

	// types of variables
	var stringWithoutValue string
	fmt.Println(stringWithoutValue) // empty string

	var intWithoutValue int
	fmt.Println(intWithoutValue) // 0

	var number float32
	number = 1.2
	fmt.Println(number)

	var b bool
	fmt.Println(b)

	var testRune rune = 'A'
	fmt.Println(testRune) // 65

	// var testByte byte = 50
	// fmt.Fprintln(testByte)
}

// go run main.go
// go build main.go
// ./main
