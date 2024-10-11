package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Input your name : ")
	name, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(name)
	}
}
