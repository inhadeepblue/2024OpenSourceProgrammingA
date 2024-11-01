package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool { // 함수화. 소수면 true 소수가 아니면 false 리턴
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
	fmt.Print("정수 입력 : ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)

	if isPrime(n) {
		fmt.Printf("%d는(은) 소수입니다", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다", n)
	}
}
