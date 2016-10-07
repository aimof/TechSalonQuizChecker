package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)
	if input == "start" {
		fmt.Printf("%s", "finish")
	} else {
		fmt.Printf("wrong input")
	}
}
