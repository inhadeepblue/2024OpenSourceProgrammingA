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
	fmt.Print("점수 입력 : ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)                  // python strip
	score, err := strconv.ParseInt(i, 10, 32) // 문자열 변수 i의 값을 정수형(32비트)으로 변환, 입력받은 값은 10진수로 처리
	var aOrNot string
	if score >= 90 {
		aOrNot = "A"
	} else {
		aOrNot = "BCDF"
	}
	fmt.Printf("%d점은 %s등급 입니다\n", score, aOrNot)
}
