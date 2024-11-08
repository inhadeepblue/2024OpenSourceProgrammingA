// print prime number between two input number
package main

import (
	"fmt"
	"log"
	"week11/keyboard"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		for j := 3; j*j <= n; j = j + 2 {
			if n%j == 0 {
				return false
			}
			//fmt.Printf("%d ", j)
		}
	}
	return true
}

func main() {
	fmt.Print("첫 번째 정수(시작 값) 입력 : ")
	n1, err := keyboard.GetInteger()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("두 번째 정수(끝 값) 입력 : ")
	n2, err := keyboard.GetInteger()
	if err != nil {
		log.Fatal(err)
	}

	for i := n1; i <= n2; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}
