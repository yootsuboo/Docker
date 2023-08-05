package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	// 文字列を単語ごとに分割する
	a := strings.Fields(s)
	// 値だけを使用する
	for _, v := range a {
		// map型の中にキーとしてvalueを入れていく。プラスした回数に応じで mapのvalueが増えていく
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
