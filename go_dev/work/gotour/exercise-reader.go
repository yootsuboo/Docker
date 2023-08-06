package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (a MyReader) Read(p []byte) (n int, e error) {
	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
}

func main() {
	reader.Validate(MyReader{})
}

// reader.Validate-------------------------------------------------------------------
//package reader // import "golang.org/x/tour/reader"
//
//import {
//	"fmt"
//	"io"
//	"os"
//)
//
//func Validate(r io.Reader) {
//	b := make([]byte, 1024, 2048)
//	i, o := 0, 0
//	for ; i < 1<<20 && o < 1<<20; i++ { // test 1MB
//		n, err := r.Read(b)
//		for i, v := range b[:n] {
//			if v != 'A' {
//				fmt.Fprintf(os.Stderr, "got byte %x at offset %v, want 'A'\n", v, o+i)
//				return
//			}
//		}
//		o += n
//		if err != nil {
//			fmt.Fprintf(os.Stderr, "read error: %v\n", err)
//			return
//		}
//	}
//	if o == 0 {
//		fmt.Fprintf(os.Stderr, "read zero bytes after %d read calls\n", i)
//		return
//	}
//	fmt.Println("OK")
//}
// end------------------------------------------------------------------------------
