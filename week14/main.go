package main

import "fmt"

func main() {
	ages := make(map[string]int)

	var name string
	var age int

	for {
		fmt.Print("What's ur name? (exit to 'q') ")
		fmt.Scanln(&name)
		if name == "q" {
			break
		}
		fmt.Print("Ur age? ")
		fmt.Scanln(&age)

		ages[name] = age
	}

	for name, age := range ages {
		fmt.Printf("%s is %d year old.\n", name, age)
	}
}
