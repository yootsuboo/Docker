package main

import (
	"fmt"
	"strings"
)

func main() {
	x := make([]int, 5, 10)
	fmt.Println(x)
	y := make([]int, 0, 10)
	fmt.Println(y)

	a := []int{1, 2, 3, 4}
	b := make([]int, 3)
	num := copy(b, a)
	fmt.Println(b, num)

	// test
	fmt.Println(strings.ToUpper("gopher"))
}
