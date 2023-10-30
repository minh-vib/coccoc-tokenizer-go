package tokenizer

/*
   #cgo LDFLAGS: -L. -ltokenizer
   #include <stdlib.h>
   #include "../tokenizer_wrapper.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

// TokenizerOption represents the options for the tokenizer.
type TokenizerOption struct {
	ptr unsafe.Pointer
}

// NewTokenizerOption creates a new TokenizerOption object.
func NewTokenizerOption() TokenizerOption {
	ptr := C.create_tokenizer_option()
	C.set_tokenizer_option_defaults(ptr)

	return TokenizerOption{ptr}
}

// SetDefaults sets the default options for the tokenizer.
func (opts TokenizerOption) SetDefaults() {
	C.set_tokenizer_option_defaults(opts.ptr)
}

// SetNoSticky sets the no sticky option.
func (opts TokenizerOption) SetNoSticky(value int) {
	C.set_no_sticky(opts.ptr, C.int(value))
}

// SetKeepPuncts sets the keep puncts option.
func (opts TokenizerOption) SetKeepPuncts(value int) {
	C.set_keep_puncts(opts.ptr, C.int(value))
}

// SetForTransforming sets the for transforming option.
func (opts TokenizerOption) SetForTransforming(value int) {
	C.set_for_transforming(opts.ptr, C.int(value))
}

// SetTokenizeOption sets the tokenize option.
func (opts TokenizerOption) SetTokenizeOption(value int) {
	C.set_tokenize_option(opts.ptr, C.int(value))
}

// SetDictPath sets the dictionary path option.
func (opts TokenizerOption) SetDictPath(path string) {
	cpath := C.CString(path)
	C.set_dict_path(opts.ptr, cpath)
}

// Destroy releases the memory allocated for the TokenizerOption object.
func (opts TokenizerOption) Destroy() {
	C.destroy_tokenizer_option(opts.ptr)
}

type Tokenizer struct {
	Options TokenizerOption
}

func NewTokenizer() Tokenizer {
	opts := NewTokenizerOption()

	return Tokenizer{Options: opts}
}

// WordTokenizer tokenizes the input text and returns a slice of strings.
func (t Tokenizer) WordTokenizer(text string) []string {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	var size C.int
	cresult := C.word_tokenizer_c(ctext, t.Options.ptr, &size)
	defer C.free_string_array(cresult, size)

	var goStrings []string
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cresult)),
		Len:  int(size),
		Cap:  int(size),
	}
	slice := *(*[]*C.char)(unsafe.Pointer(&header))
	for _, s := range slice {
		goStrings = append(goStrings, C.GoString(s))
	}

	return goStrings
}
