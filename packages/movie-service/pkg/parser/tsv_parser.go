package parser

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// ParseCompressedMovies parses the IMDb title.basics.tsv.gz file and returns movie data
func ParseCompressedMovies(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	csvReader := csv.NewReader(gzipReader)
	csvReader.Comma = '\t' // TSV uses tabs
	csvReader.LazyQuotes = true

	_, err = csvReader.Read()
	if err != nil {
		return err
	}

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		fmt.Println(record)
	}

	return nil
}
