package main

import "fmt"

func main() {
	fmt.Println(" Sup bitches.")

	foo()

	fmt.Println("continued")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("not sure what foo is")
}
