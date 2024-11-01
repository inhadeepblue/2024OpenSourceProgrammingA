package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("정수 입력 : ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)

	var isPrime bool = true

	if n < 2 {
		isPrime = false
	} else if n == 2 {
		isPrime = true
	} else if n%2 == 0 {
		isPrime = false
	} else {
		for j := 3; j*j <= n; j = j + 2 {
			if n%j == 0 {
				isPrime = false
				break
			}
			fmt.Printf("%d ", j)
		}
	}
	if isPrime {
		fmt.Printf("%d는(은) 소수입니다", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다", n)
	}
}
