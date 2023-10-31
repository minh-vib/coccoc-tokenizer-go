package tokenizer

/*
   #cgo LDFLAGS: -L. -lcoccoc-tokenizer-go
   #include <stdlib.h>
   #include "./wrapper/tokenizer_wrapper.h"
*/
import "C"
import (
	"bufio"
	"os"
	"reflect"
	"strings"
	"unsafe"
)

const (
	STOP_WORD_DEFAULT = 0
	STOP_WORD_CUSTOM  = 1

	stopWordFileName = "vi-stopwords.txt"
)

// TokenizerOption represents the options for the tokenizer.
type TokenizerOption struct {
	ptr          unsafe.Pointer
	stopWordData map[string]bool
}

// NewTokenizerOption creates a new TokenizerOption object.
func NewTokenizerOption() TokenizerOption {
	ptr := C.create_tokenizer_option()
	C.set_tokenizer_option_defaults(ptr)
	options := TokenizerOption{ptr: ptr}
	options.SetStopWordType(STOP_WORD_DEFAULT)

	return options
}

// SetDefaults sets the default options for the tokenizer.
func (opts *TokenizerOption) SetDefaults() {
	C.set_tokenizer_option_defaults(opts.ptr)
}

// SetNoSticky sets the no sticky option.
func (opts *TokenizerOption) SetNoSticky(value int) {
	if value > 1 {
		value = 1
	} else if value < 0 {
		value = 0
	}
	C.set_no_sticky(opts.ptr, C.int(value))
}

// SetKeepPuncts sets the keep puncts option.
func (opts *TokenizerOption) SetKeepPuncts(value int) {
	if value > 1 {
		value = 1
	} else if value < 0 {
		value = 0
	}
	C.set_keep_puncts(opts.ptr, C.int(value))
}

// SetForTransforming sets the for transforming option.
func (opts *TokenizerOption) SetForTransforming(value int) {
	if value > 1 {
		value = 1
	} else if value < 0 {
		value = 0
	}
	C.set_for_transforming(opts.ptr, C.int(value))
}

// SetTokenizeOption sets the tokenize option.
func (opts *TokenizerOption) SetTokenizeOption(value int) {
	C.set_tokenize_option(opts.ptr, C.int(value))
}

// SetDictPath sets the dictionary path option.
func (opts *TokenizerOption) SetDictPath(path string) {
	cpath := C.CString(path)
	C.set_dict_path(opts.ptr, cpath)
}

func (opts *TokenizerOption) SetStopWordType(value int) {
	opts.stopWordData = make(map[string]bool)
	if value == STOP_WORD_DEFAULT {
		fileName := "./dicts/" + stopWordFileName
		file, err := os.Open(fileName)
		if err != nil {
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			word := strings.ToLower(strings.TrimSpace(scanner.Text()))
			opts.stopWordData[word] = true
		}
	}
}

// Destroy releases the memory allocated for the TokenizerOption object.
func (opts *TokenizerOption) Destroy() {
	C.destroy_tokenizer_option(opts.ptr)
}

type Tokenizer struct {
	Options TokenizerOption
}

func NewTokenizer() Tokenizer {
	opts := NewTokenizerOption()

	return Tokenizer{Options: opts}
}

func (tk *Tokenizer) AddStopWords(words []string) {
	for _, word := range words {
		word = strings.ToLower(strings.TrimSpace(word))
		tk.Options.stopWordData[word] = true
	}
}

// WordTokenizer tokenizes the input text and returns a slice of strings.
func (tk *Tokenizer) WordTokenizer(text string) []string {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	var size C.int
	cresult := C.word_tokenizer_c(ctext, tk.Options.ptr, &size)
	defer C.free_string_array(cresult, size)

	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cresult)),
		Len:  int(size),
		Cap:  int(size),
	}
	slice := *(*[]*C.char)(unsafe.Pointer(&header))

	var words []string
	for _, s := range slice {
		word := strings.ToLower(strings.TrimSpace(C.GoString(s)))
		if ok, _ := tk.Options.stopWordData[word]; !ok {
			words = append(words, word)
		}
	}

	return words
}
