package main

import "fmt"

//----interface-values--------------------------------
// 今のところ全然理解できていない 2023/08/06
//type I interface {
//	M()
//}
//
//type T struct {
//	S string
//}
//
//func (t *T) M() {
//	fmt.Println(t.S)
//}
//
//type F float64
//
//func (f F) M() {
//	fmt.Println(f)
//}
//
//func main() {
//  //このIの意味がわからない
//  // I = interface
//  // Iを呼び出すと基底型のM()メソッドが呼び出される
//	var i I
//
//	i = &T{"Hello"}
//	describe(i)
//	i.M()
//}
//
//func describe(i I) {
//	fmt.Printf("(%v, %T)\n", i, i)
//}
//----end-----------------------------------------------

//----interface-values-with-nil-------------------------

// nilインターフェースの値は具体的な型を持たない
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I
	// 具体的な値としてのnilを保持している場合は、非nilとなる??
	// 以下は非nilなのか??
	fmt.Println("i= ", i)

	// Tに値が設定されていないため、nilレシーバとして呼び出される
	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
