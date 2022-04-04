package main

import (
	"fmt"

	"01_hello_world/mypackage"
)

func main() {								// main() cant take any parameter
	fmt.Println("Hello, World! this is Println speaking!")
	sayHelloWorld("Hello, World! this is function sayHelloWorld speaking!") // function calls inside main package
	mypackage.PrintHello()   			// package calls inside main package
}

func sayHelloWorld(whatToSay string) { 		// function with parameter whatToSay in string
	fmt.Println(whatToSay)
}