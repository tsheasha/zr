package data

import (
	"fmt"
	"strings"

	"github.com/dgryski/go-metro"
	"github.com/rekki/go-query/util/analyzer"
	"github.com/rekki/go-query/util/norm"
	"github.com/rekki/go-query/util/tokenize"
)

var DefaultNormalizer = []norm.Normalizer{
	norm.NewCustom(func(s string) string {
		var sb strings.Builder
		hadSpace := false
		for _, c := range s {
			if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9') {
				if 'A' <= c && c <= 'Z' {
					c += 32
				}
				sb.WriteRune(c)
				hadSpace = false
			} else {
				if hadSpace {
					continue
				} else {
					sb.WriteRune(' ')
					hadSpace = true
				}
			}
		}
		return sb.String()
	}),
}

var trimmer = tokenize.NewCustom(func(in []string) []string {
	for i := range in {

		first := in[i][0]

		h := metro.Hash64Str(in[i], 0)

		// 65k per starting character
		// so overall 65_000 * 36, or about 2.5 mil files

		in[i] = fmt.Sprintf("%x_%c", h&0x000000000000FFFF, first)
	}
	return in
})

var DefaultSearchTokenizer = []tokenize.Tokenizer{
	tokenize.NewWhitespace(),
	trimmer,
}

var DefaultIndexTokenizer = []tokenize.Tokenizer{
	tokenize.NewWhitespace(),
	trimmer,
	tokenize.NewUnique(),
}

var DefaultAnalyzer = analyzer.NewAnalyzer(DefaultNormalizer, DefaultSearchTokenizer, DefaultIndexTokenizer)
