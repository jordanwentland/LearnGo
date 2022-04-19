package main

import (
	"fmt"
)

func squareAreaCalc(x, y float64) float64 {
	return (x * y)
}

func main() {
	var a, b, c float64
	fmt.Println("===AREA OF SQUARE CALC USING FUNCTIONS===")
	fmt.Println("Please enter the length of the square: ")
	fmt.Scan(&a)
	fmt.Println("Please enter the width of the square: ")
	fmt.Scan(&b)
	c = squareAreaCalc(a, b)
	fmt.Println("the area of your square is", c, "units squared!")
}
