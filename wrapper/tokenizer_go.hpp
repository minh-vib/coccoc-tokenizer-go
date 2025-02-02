#include "../tokenizer/tokenizer.hpp"
#include "../tokenizer/config.h"

struct tokenizer_option
{
    bool no_sticky;
    bool keep_original;
    bool keep_puncts;
    bool for_transforming;
    int tokenize_option;
    const char *dict_path;

    tokenizer_option()
        : no_sticky(true),
          keep_original(false),
          keep_puncts(false),
          for_transforming(false),
          tokenize_option(Tokenizer::TOKENIZE_NORMAL),
          dict_path(DICT_PATH)
    {
    }
};

std::vector< std::string > word_tokenizer(const std::string &text, tokenizer_option &opts) {
    std::vector< std::string > result;

    if (text.length() == 0 || 0 > Tokenizer::instance().initialize(opts.dict_path, !opts.no_sticky))
    {
        return result;
    }

    std::vector< FullToken > res = opts.keep_original ?
        Tokenizer::instance().segment_original(text, opts.tokenize_option) :
        Tokenizer::instance().segment(text, opts.for_transforming, opts.tokenize_option, opts.keep_puncts);

    std::string word;
    for (size_t i = 0; i < res.size(); ++i)
    {
        word = res[i].text;
        if (opts.keep_original) {
            std::replace(word.begin(), word.end(), '_', ' ');
            std::replace(word.begin(), word.end(), '~', '_');
        }
        result.push_back(word);
    }
    return result;
}
