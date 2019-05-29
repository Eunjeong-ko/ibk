package main

import "fmt"

func main() {
	testparam("gogo")
}

// param name 먼저, type 뒤에!
func testparam(a string) {
	fmt.Println(a)
}