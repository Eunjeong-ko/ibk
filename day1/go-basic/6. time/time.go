package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now().Format("2006-01-02 15:04:05") // yyyymmhhddss 숫자는 정해져있음
	p(now)
}