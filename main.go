package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	foo()
	fmt.Println("after foo")

	for i := 0; i < 100; i++ {
		if i %2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I'm in foo.")
}