package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	//var month int = int(now.Month())
	fmt.Printf("오늘은 %d년 %d월 %d일 입니다\n", now.Year(), int(now.Month()), now.Day())
	fmt.Printf("지금은 시각은 %d시 %d분 %d초 입니다\n", now.Hour(), now.Minute(), now.Second())
	fmt.Println(now.Month())
}
