// Task1: Swap 2 numbers
// In this task, user is asked to enter two numbers and program will swap two numbers without using third variable

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Printf("a = %d, b = %d\n", a, b)

	a, b = b, a
	fmt.Printf("a = %d, b = %d\n", a, b)
}
