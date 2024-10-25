package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(6) + 1 // 1 ~ 3
	fmt.Printf("%d\n", target)
}
