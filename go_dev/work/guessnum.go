package main 

import (
    "fmt"
    "strconv"
)

func main() {
    answer := 4
    fmt.Print("数当てゲームです。1以上10以下の整数を(半角で)入力してください: ")
    var inp string
    fmt.Scanln(&inp)
    num, err := strconv.Atoi(inp)
    if err != nil || num <1 || 10 < num {
        fmt.Println("1以上10以下の整数ではないので、外れです。")
    } else if num == answer {
        fmt.Println("ビンゴ")
    } else {
        fmt.Println("残念でした。ハズレです。")
    }
}
