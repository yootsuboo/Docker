package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	pow := make([]int, 10)
	for i := range pow {
		// ここの書き方の意味がよくわからない 2023/08/05
		pow[i] = 1 << uint(i) // == 2**i
		fmt.Println(pow[i])
	}
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}
