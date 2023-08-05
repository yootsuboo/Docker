package main

import "fmt"

// -----map----------------------------------------------
//type Vertex struct {
//	Lat, Long float64
//}
//
//var m = map[string]Vertex{
//	"Bell Labs": {40.68433, -74.39967},
//	"Google":    {37.42202, -122.08400},
//}
//
//func main() {
//	// 一部だけを取得する方法がわからない 2023/08/05
//	// fmt.Println(m["Bell Labs"[Let]])
//	fmt.Println(m)
//}
// ----end-----------------------------------------------

// ----mutating-maps-------------------------------------
func main() {
	// keyがstringで、valueがint型となる
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// keyに対して複数の値が設定できるわけではない
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// keyに対して一つしか値がないので選択して削除などはないし、必要ない
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// 宣言されていないことが条件の書き方なので、新しい変数で実施する
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
