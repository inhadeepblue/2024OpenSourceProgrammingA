package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var month int = int(now.Month())
	fmt.Println(month)
}
