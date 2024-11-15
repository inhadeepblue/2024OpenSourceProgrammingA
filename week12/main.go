package main

import (
	"fmt"
	"reflect"
)

func main() {
	var gpa [5]float64 = [5]float64{3.5, 4.1, 4.5, 3.9, 4.23} // array
	gpa_slice := gpa[1:4]                                     // slicing
	fmt.Println(gpa_slice, reflect.TypeOf(gpa_slice))
	fmt.Println(gpa, reflect.TypeOf(gpa))
}
