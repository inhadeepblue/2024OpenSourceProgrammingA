package main

import "fmt"

type student struct {
	id   int
	name string
	gpa  float32
}

func main() {
	var student1 student
	student1.id = 20241234
	student1.name = "King sooa"
	student1.gpa = 4.5
	fmt.Println(student1.gpa)
	var student2 student
	student2.id = 20244321
	student2.name = "Choi joonhoon"
	student2.gpa = 4.7
	fmt.Println(student2.id)
}
