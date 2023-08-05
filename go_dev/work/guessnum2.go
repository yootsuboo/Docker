package main

import (
    "fmt"
    "strconv"
    "math/rand"
    "time"
)
func main() {
    answer := getTheNumber();

    for count :=1; ; count++ {
        printPrompt(count)

        if num, err := readUserAnswer(count);
        err != nil || num < 1 || 10 < num {
            fmt.Println("1以上10以下の整数ではないので、ハズレです。")
        } else if num != answer {
            fmt.Println("残念でした。ハズレです")
        } else {
            printSuccessMessage(count)
            break
        }
    }
}

func getTheNumber() int {
    rand.Seed(time.Now().UnixNano())
    num := rand.Intn(10)+1
    return num
}

func printPrompt(count int) {
    if count == 1 {
        fmt.Print("数当てゲームです。")
    }
    fmt.Printf("1以上10以下の整数(半角で)入力してください(%v回目) : ", count)
}

func readUserAnswer(count int) (int, error) {
    var inp string
    fmt.Scanln(&inp)
    return strconv.Atoi(inp)
}

func printSuccessMessage(count int) {
    if count == 1 {
        fmt.Printf("ビンゴ! おめでとうございます。一発で当たりましたね。素晴らしい!\n")
    } else {
        adverb := "" 
        if 7 < count {
            adverb = "ヤット"
        }
        fmt.Printf("おめでとうございます。%v回目で%s当たりましたね。\n", count, adverb)
    }
}
