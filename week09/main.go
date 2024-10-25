package main

import (
	"fmt"
)

func main() {
	var result string
	result = fmt.Sprintf("%0.1f 나누기 %0.1f(은)는 %0.2f입니다\n", 1.0, 3.0, 1.0/3.0)
	fmt.Print(result)

	i := 1
	for i <= 10 {
		fmt.Printf("%d\n", i)
		i++
	}
}
