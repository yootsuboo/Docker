// const は定数なのであとから値を変更することができない
// const の値を var に代入することは可能

package main

import "fmt"

// 型を指定することができない -> int 型は指定することが可能
// const x int64 = 10
// 型付きの定数を宣言した場合には、他の型に代入することができない
const x int = 10

var y int = x
var z float64 = x
var d byte = x

const (
    idkey = "id"
    namekey = "name"
)

const z = 20 * 10

func main() {
    // const y = "hello"

    fmt.Println(x)
    // fmt.Println(y)

    // const の値を変更することはできない
    // x = x + 10
    y := "bye"

    fmt.Println(x)
    fmt.Println(y)
}


