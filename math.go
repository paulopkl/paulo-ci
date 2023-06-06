package main

import "fmt"

func main() {
	sum := sum(13, 14)
	minus := Minus(15, 8)

	fmt.Println(sum + minus)
}

func sum(a int, b int) int {
	return a + b
}

func Minus(a int, b int) int {
	return a - b
}
