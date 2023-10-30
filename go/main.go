package main

import (
	"fmt"

	"coccoc-tokenizer-go/tokenizer"
)

func main() {
	opts := tokenizer.NewTokenizerOption()

	text := "toisongohanoi, tôi đăng ký trên Thegioididong.vn"
	result := tokenizer.WordTokenizer(text, opts)

	fmt.Println(result)

	opts.Destroy()
}
