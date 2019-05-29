package main

import "fmt"

func main() {
	fmt.Println(testReturn("1"))
}

func testReturn(a string) (string, string) {
	fmt.Println(a)
	return "2", "3"
}