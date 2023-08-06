package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 最大256
type IPAddr [4]byte

func (ip IPAddr) String() string {
	// ipのスライスの個数に応じで[ len(ip) ]サイズを指定している
	s := make([]string, len(ip))
	for i, val := range ip {
		s[i] = strconv.Itoa(int(val))
	}
	return strings.Join(s, ".")
}

func main() {
	// map はpythonでの辞書のようなもの キー:バリューを指定する
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
