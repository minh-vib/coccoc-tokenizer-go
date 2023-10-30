package main

import (
	"fmt"

	"coccoc-tokenizer-go/tokenizer"
)

func main() {
	token := tokenizer.NewTokenizer()
	token.Options.SetDictPath("/usr/local/share/tokenizer/dicts")

	text := "toisongohanoi, tôi đăng ký trên Thegioididong.vn"
	result := token.WordTokenizer(text)

	fmt.Println(result)

	token.Options.Destroy()
}
