package parser

import (
	"compress/gzip"
	"encoding/csv"
	"io"
	"movie-service/internal/models"
	"os"
)

// ParseCompressedMovies parses the IMDb title.basics.tsv.gz file and returns movie data
func ParseCompressedMovies(filePath string) ([]*models.Movie, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	csvReader := csv.NewReader(gzipReader)
	csvReader.Comma = '\t' // TSV uses tabs
	csvReader.LazyQuotes = true

	_, err = csvReader.Read()
	if err != nil {
		return nil, err
	}

	var movies []*models.Movie
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		if record[1] == "movie" {
			movies = append(movies, models.NewMovie(record))
		}
	}

	return movies, nil
}
