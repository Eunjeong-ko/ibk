package main

import "fmt"

func main() {
	var a = "this is a"
	// Println 대문자 시작 -> public (소문자인경우, private)
	fmt.Println(a)

	// 변수 초기화 ":"
	b := "this is b"
	b = "this is c"
	fmt.Println(b)
}