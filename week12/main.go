package main

import (
	"fmt"
	"time"
)

func main() {
	// var dates [3]time.Time
	// dates[1] = time.Unix(1947920001, 0)
	// fmt.Println(dates[1])

	// var dates [3]time.Time = [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1947920001, 0)}
	// fmt.Println(dates[0], dates[1], dates[2])

	// dates := [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1947920001, 0)}
	// fmt.Println(dates[0], dates[1], dates[2])

	// dates := [3]time.Time{
	// 	time.Unix(0, 0),
	// 	time.Unix(1, 0),
	// 	time.Unix(1947920001, 0)}
	// fmt.Println(dates[0], dates[1], dates[2])

	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(1947920001, 0), // need comma
	}
	//fmt.Println(dates[0], dates[1], dates[2])  // array elements
	//fmt.Println(dates)  // array
	//fmt.Printf("%#v\n", dates)  // array literal

	// for i := 0; i < len(dates); i++ { // safe
	// 	//for i := 0; i <= 3; i++ {  // panic: runtime error: index out of range
	// 	fmt.Println(i, dates[i])
	// }

	for i, v := range dates { // safe
		fmt.Println(i, v)
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
