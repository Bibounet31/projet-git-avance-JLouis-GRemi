package main

import "fmt"

// main func, calling others
func main() {
	fmt.Println("Hello World")
	Counttoten()
	Dontcounttoten()
}

// a function that count from 0 to 10
func Counttoten() {
	for i := 0; i < 11; i++ {
		fmt.Println(i)
	}
}

// a function that doesn't count from 0 to 10
func Dontcounttoten() {
	fmt.Println("a")
}
