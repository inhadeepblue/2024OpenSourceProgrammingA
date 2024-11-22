package main

import (
	"fmt"
	"os"
	"reflect"
)

// func test(strs string) {
func test(strs ...string) {
	fmt.Println(strs, reflect.TypeOf(strs))
}

func main() {
	//fmt.Println(os.Args)
	slice := os.Args[1:]
	fmt.Println(slice[1])
	fmt.Printf("%T\n", slice[1])
	slice = append(slice, "i", "n", "h", "a")
	fmt.Println(slice, len(slice))
	test("123", "456")
	test("123")
	test()
	test("123", "456", "0999")
}
