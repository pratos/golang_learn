package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 2, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	z := f
	fmt.Printf("%v %v %0.2f \n", x, y, z)
}
