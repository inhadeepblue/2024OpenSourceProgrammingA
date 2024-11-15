package main

import (
	"fmt"

	"github.com/headfirstgo/keyboard"
)

func main() {
	var gpa [3]float64

	for i := 0; i < len(gpa); i++ {
		fmt.Print("Input float number : ")
		gpa[i], _ = keyboard.GetFloat() // go get github.com/headfirstgo/keyboard
	}
	for index, value := range gpa {
		fmt.Printf("%d %f\n", index, value)
	}
}

/*
import random

numbers = [random.randint(1, 100) for _ in range(10)]
print(numbers)


#for i in range(len(numbers)):
#for i in range(11):  # Runtime error
#	print(numbers[i], end=' ')

for number in numbers:  # safe
	print(number, end=' ')
*/
