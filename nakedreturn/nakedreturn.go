package main

import (
	"fmt"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return //A naked return here, without any arguments: It would return both x and y
}

func main() {
	fmt.Println(split(18))
}
