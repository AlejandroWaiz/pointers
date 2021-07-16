package main

import "fmt"

func main() {

	foo := "I am the original"

	bar := *&foo

	bar = "No! I am the original."

	fmt.Println(foo)
	fmt.Println(bar)

}
