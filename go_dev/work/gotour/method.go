package main

import (
	"fmt"
	"math"
)

//----methods------------------------------------------------------
//type Vertex struct {
//	X, Y float64
//}
//
//// Absというメソッドの中にvという名前のVertex型のレシーバがある
//func (v Vertex) Abs() float64 {
//	// 3*3 + 4*4 の平方根
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//func main() {
//	v := Vertex{3, 4}
//	fmt.Println(v.Abs())
//}
//----end---------------------------------------------------------

//----methods-continued-------------------------------------------
//type MyFloat float64
//
//func (f MyFloat) Abs() float64 {
//	if f < 0 {
//		return float64(-f)
//	}
//	return float64(f)
//}
//
//func main() {
//	// -math.Sqrt2 の書き方の意味がわからない 2を3にしたらエラーになった 2023/08/05
//	// 2以外のときには()が必要??
//	// -math にすることでSqrtの値が「-」になるようです
//	f := MyFloat(-math.Sqrt(3))
//	fmt.Println(f.Abs())
//}
//----end---------------------------------------------------------

// ----methods-pointers--------------------------------------------
//type Vertex struct {
//	X, Y float64
//}
//
//func (v Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//func (v *Vertex) Scale(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//
//func main() {
//	// *Vertexとしてポインタレシーバにしないと値が更新されない
//	v := Vertex{3, 4}
//	v.Scale(10)
//	fmt.Println(v.Abs())
//}
//----end---------------------------------------------------------

//----methods-pointers-explained----------------------------------
//type Vertex struct {
//	X, Y float64
//}
//
//func Abs(v Vertex) float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//func Scale(v *Vertex, f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//
//func main() {
//	v := Vertex{3, 4}
//	// *Vertexではなくなった場合、ポインタではなくなるので`&`が不要になる
//	Scale(&v, 10)
//	fmt.Println(Abs(v))
//}
//----end---------------------------------------------------------

// ----indirection------------------------------------------------
//type Vertex struct {
//	X, Y float64
//}
//
//func (v *Vertex) Scale(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//
//func ScaleFunc(v *Vertex, f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//
//func main() {
//	v := Vertex{3, 4}
//	// vは変数でありポイントではないが、ポインタレシーバが自動的に呼び出される
//	// (&v).Scale(2)として解釈されている
//	v.Scale(2)
//	ScaleFunc(&v, 10)
//
//	p := &Vertex{4, 3}
//	p.Scale(3)
//	ScaleFunc(p, 8)
//
//	fmt.Println(v, p)
//}
//----end-----------------------------------------------------------

// ----indirection-values--------------------------------------------
//type Vertex struct {
//	X, Y float64
//}
//
//func (v Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//func AbsFunc(v Vertex) float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//func main() {
//	v := Vertex{3, 4}
//	fmt.Println(v.Abs())
//	// AbsFunc(&v)とするとエラーとなる
//	// 変数の引数を取る関数は、特定の型の変数を取る必要がある
//	fmt.Println(AbsFunc(v))
//
//	// &Vertexなので変数レシーバとなっている
//	p := &Vertex{4, 3}
//	// p.Abs() は (*p).Abs() として解釈されている
//	fmt.Println(p.Abs())
//	fmt.Println(AbsFunc(*p))
//}
//----end-----------------------------------------------------------

// ポインタレシーバを使う理由
// メソッドがレシーバが指す先の変数を変更するため
// メソッド呼び出し毎に変数のコピーを避けるため

// 値レシーバとポインタレシーバの違いとは?? 2023/08/06

// ----methods-with-pointer-receivers--------------------------------
//type Vertex struct {
//	X, Y float64
//}
//
//func (v *Vertex) Scale(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//
//func (v *Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//func main() {
//	v := &Vertex{3, 4}
//	// %+ とすることでstruct内の情報を追加してフィールド名がkeyとして表示されるようになる
//	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
//
//	v.Scale(5)
//	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
//}
//----end-------------------------------------------------------------

// ----interfaces------------------------------------------------------
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	fmt.Println("f: ", a.Abs())
	a = &v

	// Error
	// a = v

	fmt.Println("v: ", a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
