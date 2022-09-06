//Task3: Area of circle inscribed in a square
// Find the shaded region by given R(radius of the circle)

package main

import (
	"fmt"
)

func main() {
	var r float32 = 10.04

	// fmt.Scanf("%f", &r)
	fmt.Println("R =", r)

	fmt.Printf("Area: %0.2f\n", area(r))
}

func area(r float32) (area float32) {
	const PI = 3.14159265

	var length_square = r * 2
	var area_square = length_square * length_square
	var area_circle = PI * r * r
	area = area_square - area_circle

	return area
}
