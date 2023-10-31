package main

import (
	"fmt"

	"github.com/minh-vib/coccoc-tokenizer-go"
)

func main() {
	tk := tokenizer.NewTokenizer()
	tk.Options.SetDictPath("/usr/local/share/tokenizer/dicts")
	//tk.Options.SetStopWordType(tokenizer.STOP_WORD_CUSTOM)
	tk.AddStopWords([]string{"đăng ký"})

	text := "toisongohanoi, tôi đăng ký trên Thegioididong.vn"
	result := tk.WordTokenizer(text)

	fmt.Println(result)

	tk.Options.Destroy()
}
