package main

import (
	"fmt"
	"strings"

	"github.com/minh-vib/coccoc-tokenizer-go"
)

func main() {
	tk := tokenizer.NewTokenizer()
	tk.Options.SetNoSticky(0)
	tk.Options.SetKeepOriginal(1)
	//tk.Options.SetDictPath("/usr/local/share/tokenizer/dicts")

	text := "Từng bước để trở thành một lập trình viên giỏi"
	result := tk.WordTokenizer(text)
	fmt.Println("\"" + strings.Join(result, "\",\"") + "\"")

	text = "toisongohanoi, tôi đăng ký trên Thegioididong.vn"
	result = tk.WordTokenizer(text)
	fmt.Println("\"" + strings.Join(result, "\",\"") + "\"")

	tk.Options.Destroy()
}
