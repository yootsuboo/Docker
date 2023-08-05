package main

import "fmt"

// ----function-values-------------------------------------------
//func compute(fn func(float64, float64) float64) float64 {
//	return fn(3, 4)
//}
//
//func main() {
//	hypot := func(x, y float64) float64 {
//		return math.Sqrt(x*x + y*y)
//	}
//	fmt.Println(hypot(5, 12))
//	fmt.Println(compute(hypot))
//  // ()を付けないと関数のIDが返されるだけになる returnが実行されない
//	fmt.Println(compute())
//	fmt.Println(compute(math.Pow))
//}
// ----end-------------------------------------------------------

// ----function-closuers-----------------------------------------
//func adder() func(int) int {
//	sum := 0
//	fmt.Println(sum)
//	return func(x int) int {
//		sum += x
//		return sum
//	}
//}
//
//func main() {
//	pos, neg := adder(), adder()
//	for i := 0; i < 10; i++ {
//		fmt.Println(pos(i), neg(-2*i))
//	}
//}
// ----end------------------------------------------------------

// ----exercise-fibonacci-closure-------------------------------

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		// ここで何をしているのかかが理解できない y = y ???
		x, y = y, x+y
		fmt.Printf("y=%d \n", y)
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		// fの関数を実行しているから()が必要
		fmt.Println(f())
	}
}
