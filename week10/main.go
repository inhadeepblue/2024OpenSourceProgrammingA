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

	// counts := 0
	var isPrime bool = true  // 가독성, 저장공간크기 개선
	for j := 2; j < n; j++ { // 2부터 입력된 수 앞까지 반복
		if n%j == 0 { // 약수면
			//counts++ // counts = counts + 1 , 더하기 연산 제거
			isPrime = false
		}
	}

	//if counts == 0 {  // 나누어 떨어지는 수가 있으면 안됨
	if isPrime { // == 비교 연산 제거
		fmt.Printf("%d는(은) 소수입니다", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다", n)
	}
}
