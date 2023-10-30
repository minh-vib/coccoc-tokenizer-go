#include "../tokenizer/tokenizer.hpp"
#include "../tokenizer/config.h"

struct tokenizer_option
{
    bool no_sticky;
    int keep_puncts;
    bool for_transforming;
    int tokenize_option;
    const char *dict_path;

    tokenizer_option()
        : no_sticky(false),
          keep_puncts(-1),
          for_transforming(false),
          tokenize_option(Tokenizer::TOKENIZE_NORMAL),
          dict_path("dicts")
    {
    }
};

std::vector< std::string > word_tokenizer(const std::string &text, tokenizer_option &opts) {
    std::vector< std::string > result;

    if (text.length() == 0 || 0 > Tokenizer::instance().initialize(opts.dict_path, !opts.no_sticky))
    {
        return result;
    }

    if (opts.keep_puncts == -1) {
        opts.keep_puncts = opts.for_transforming;
    }
    std::vector< FullToken > res = Tokenizer::instance().segment_original(text, opts.tokenize_option);

    size_t i = 0;

    for (/* void */; i < res.size(); ++i)
    {
        size_t punct_start = (i > 0) ? res[i - 1].original_end : 0;
        size_t punct_len = res[i].original_start - punct_start;
        if (punct_len > 0)
        {
            result.push_back(text.substr(punct_start, punct_len));
        }

        result.push_back(res[i].text);
    }

    size_t punct_start = (i > 0) ? res[i - 1].original_end : 0;
    size_t punct_len = text.size() - punct_start;
    if (punct_len > 0)
    {
        result.push_back(text.substr(punct_start, punct_len));
    }

    return result;
}
