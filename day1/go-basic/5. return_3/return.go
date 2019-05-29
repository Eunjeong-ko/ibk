package main

import "fmt"

func main() {
	// var1, var2 := testReturn()
	// var2를 나중에 사용하지 않으면 error
	// 사용하지 않는 변수를 "_"로 표시

	var1, _ := testReturn()
	fmt.Println(var1)
	//fmt.Println(var2)
}

func testReturn() (string, string) {
	return "3", "2"
}