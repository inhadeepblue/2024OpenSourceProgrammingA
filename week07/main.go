package main

import (
	"fmt"
	"strings"
)

func main() {
	var army string = "우리는 ?가와 ?민에 충성을 다하는 대한민? 육군이다"
	armyFixed := strings.NewReplacer("?", "국")
	fmt.Println(army)
	fmt.Println(armyFixed.Replace(army))
}
