package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// ここの書き方が理解できていない
	pic := make([][]uint8, dy)
	// ここのfor文と次のfor文も理解できていない
	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			pin[y][x] = uint8((x*x + y*y) / 2 >> 2)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
