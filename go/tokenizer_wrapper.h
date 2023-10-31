#ifndef TOKENIZER_WRAPPER_H
#define TOKENIZER_WRAPPER_H

struct tokenizer_option_c {
    int no_sticky;
    int keep_puncts;
    int for_transforming;
    int tokenize_option;
    const char *dict_path;
};

#ifdef __cplusplus
extern "C" {
#endif

void* create_tokenizer_option();
void destroy_tokenizer_option(void* ptr);
void set_tokenizer_option_defaults(void* ptr);
void set_no_sticky(void* ptr, int value);
void set_keep_puncts(void* ptr, int value);
void set_for_transforming(void* ptr, int value);
void set_tokenize_option(void* ptr, int value);
void set_dict_path(void* ptr, const char* path);
char** word_tokenizer_c(const char* text, void* opts, int* size);
void free_string_array(char** array, int size);

#ifdef __cplusplus
}
#endif

#endif // TOKENIZER_WRAPPER_H
