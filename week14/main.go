package main

import "fmt"

type visitor struct {
	age  int
	cost int
}

func calculateCost(visitors []visitor) int {
	totalCost := 0
	for _, v := range visitors {
		totalCost = totalCost + v.cost
	}
	return totalCost
}

func main() {
	var numVisitors int
	fmt.Println("How many visitors? ")
	fmt.Scanln(&numVisitors)

	var vs []visitor // 구조체 슬라이스 변수
	vs = make([]visitor, numVisitors)

	for i := 0; i < numVisitors; i++ {
		var age int
		fmt.Print("Input age : ")
		fmt.Scan(&age)

		switch {
		case age < 12:
			vs[i] = visitor{age: age, cost: 5000}
		case age >= 12 && age < 65:
			vs[i] = visitor{age: age, cost: 10000}
		default: // over 65 years old
			vs[i] = visitor{age: age, cost: 7000}
		}
	}

	fmt.Printf("Total price is %dwon.", calculateCost(vs))
}
