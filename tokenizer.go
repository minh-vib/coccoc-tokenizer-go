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
	DICT_PATH = "/usr/local/share/tokenizer/dicts"

	STOP_WORD_DEFAULT = 0
	STOP_WORD_CUSTOM  = 1

	stopWordFileName = "vi-stopwords.txt"
)

type Tokenizer struct {
	Options TokenizerOption
}

// TokenizerOption represents the options for the tokenizer.
type TokenizerOption struct {
	ptr          unsafe.Pointer
	dictPath     string
	stopWordType int
	stopWordData map[string]bool
}

// NewTokenizerOption creates a new TokenizerOption object.
func NewTokenizerOption() TokenizerOption {
	ptr := C.create_tokenizer_option()
	C.set_tokenizer_option_defaults(ptr)

	options := TokenizerOption{ptr: ptr}
	options.dictPath = DICT_PATH
	options.SetStopWordType(STOP_WORD_DEFAULT)

	return options
}

// SetDefaults sets the default options for the tokenizer.
func (opt *TokenizerOption) SetDefaults() {
	C.set_tokenizer_option_defaults(opt.ptr)
}

// SetNoSticky sets the no sticky option.
func (opt *TokenizerOption) SetNoSticky(value int) {
	if value > 1 {
		value = 1
	} else if value < 0 {
		value = 0
	}
	C.set_no_sticky(opt.ptr, C.int(value))
}

// SetKeepPuncts sets the keep puncts option.
func (opt *TokenizerOption) SetKeepPuncts(value int) {
	if value > 1 {
		value = 1
	} else if value < 0 {
		value = 0
	}
	C.set_keep_puncts(opt.ptr, C.int(value))
}

// SetForTransforming sets the for transforming option.
func (opt *TokenizerOption) SetForTransforming(value int) {
	if value > 1 {
		value = 1
	} else if value < 0 {
		value = 0
	}
	C.set_for_transforming(opt.ptr, C.int(value))
}

// SetTokenizeOption sets the tokenize option.
func (opt *TokenizerOption) SetTokenizeOption(value int) {
	C.set_tokenize_option(opt.ptr, C.int(value))
}

// SetDictPath sets the dictionary path option.
func (opt *TokenizerOption) SetDictPath(path string) {
	cpath := C.CString(path)
	C.set_dict_path(opt.ptr, cpath)

	opt.dictPath = path
	if len(opt.stopWordData) == 0 {
		opt.SetStopWordType(opt.stopWordType)
	}
}

func (opt *TokenizerOption) SetStopWordType(value int) {
	opt.stopWordType = value
	opt.stopWordData = make(map[string]bool)
	if value == STOP_WORD_DEFAULT {
		fileName := opt.dictPath + "/" + stopWordFileName
		file, err := os.Open(fileName)
		if err != nil {
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			word := strings.ToLower(strings.TrimSpace(scanner.Text()))
			opt.stopWordData[word] = true
		}
	}
}

// Destroy releases the memory allocated for the TokenizerOption object.
func (opt *TokenizerOption) Destroy() {
	C.destroy_tokenizer_option(opt.ptr)
}

func NewTokenizer() Tokenizer {
	opt := NewTokenizerOption()
	return Tokenizer{Options: opt}
}

func (tkn *Tokenizer) AddStopWords(words []string) {
	for _, word := range words {
		word = strings.ToLower(strings.TrimSpace(word))
		tkn.Options.stopWordData[word] = true
	}
}

// WordTokenizer tokenizes the input text and returns a slice of strings.
func (tkn *Tokenizer) WordTokenizer(text string) []string {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	var size C.int
	cresult := C.word_tokenizer_c(ctext, tkn.Options.ptr, &size)
	defer C.free_string_array(cresult, size)

	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cresult)),
		Len:  int(size),
		Cap:  int(size),
	}
	slice := *(*[]*C.char)(unsafe.Pointer(&header))

	var goStrings []string
	for _, s := range slice {
		goString := strings.ToLower(strings.TrimSpace(C.GoString(s)))
		if ok, _ := tkn.Options.stopWordData[goString]; !ok {
			goStrings = append(goStrings, goString)
		}
	}

	return goStrings
}
