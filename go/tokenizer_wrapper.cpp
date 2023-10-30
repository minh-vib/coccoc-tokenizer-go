#include "tokenizer_wrapper.h"
#include "tokenizer_go.hpp"

void* create_tokenizer_option() {
    return new tokenizer_option();
}

void destroy_tokenizer_option(void* ptr) {
    delete static_cast<tokenizer_option*>(ptr);
}

void set_tokenizer_option_defaults(void* ptr) {
    tokenizer_option* opts = static_cast<tokenizer_option*>(ptr);
    opts->no_sticky = true;
    opts->keep_puncts = -1;
    opts->for_transforming = false;
    opts->tokenize_option = Tokenizer::TOKENIZE_NORMAL;
    opts->dict_path = DICT_PATH;
}

void set_no_sticky(void* ptr, int value) {
    tokenizer_option* opts = static_cast<tokenizer_option*>(ptr);
    opts->no_sticky = value;
}

void set_keep_puncts(void* ptr, int value) {
    tokenizer_option* opts = static_cast<tokenizer_option*>(ptr);
    opts->keep_puncts = value;
}

void set_for_transforming(void* ptr, int value) {
    tokenizer_option* opts = static_cast<tokenizer_option*>(ptr);
    opts->for_transforming = value;
}

void set_tokenize_option(void* ptr, int value) {
    tokenizer_option* opts = static_cast<tokenizer_option*>(ptr);
    opts->tokenize_option = value;
}

void set_dict_path(void* ptr, const char* path) {
    tokenizer_option* opts = static_cast<tokenizer_option*>(ptr);
    opts->dict_path = path;
}

char** word_tokenizer_c(const char* text, void* opts, int* size) {
    tokenizer_option* options = static_cast<tokenizer_option*>(opts);
    std::vector<std::string> result = word_tokenizer(text, *options);
    *size = result.size();

    char** array = (char**)malloc(*size * sizeof(char*));
    for (int i = 0; i < *size; ++i) {
        array[i] = strdup(result[i].c_str());
    }

    return array;
}

void free_string_array(char** array, int size) {
    for (int i = 0; i < size; ++i) {
        free(array[i]);
    }
    free(array);
}
