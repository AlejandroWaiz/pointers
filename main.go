package main

import "fmt"

func main() {

	//Here 'foo' is var declaration, that means, the space's memory and "I am the original" is the value
	//holded on that space.
	foo := "Hurry up! Save me, a clon wants to kill me!"

	//So here we declare a var that POINTS to the value in the memory that is called 'foo'.
	bar := &foo

	//And now, to 'dereference' it, that means, use the actual value in the memory and NOT the memory itsealf
	//we use the * operator to change the value
	*bar = "Now i am the original. *dramatic sounds*"

	//Finally, to make sure that this works, we just have to know either if the original or the clon
	//are still alive
	fmt.Println(foo)

}
