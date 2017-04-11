package main

import "fmt"

//much larger value than int64
const myConst = 9223372036854775808543522345

// We add int64, limiting the number program will no longer compile
// Error: github.com/golang_learn/numericonstant/numericonstant.go:9: constant 9223372036854775808543522345 overflows int64
const myConst1 int64 = 9223372036854775808543522345

func main() {
	fmt.Println("Compiling...")
}
