package parser_test

import (
	"movie-service/pkg/parser"
	"testing"
)

func TestTSVParser(t *testing.T) {
	parser.ParseCompressedMovies("/home/kang/Documents/imdbMovies/data/title.basics.tsv.gz")
}
