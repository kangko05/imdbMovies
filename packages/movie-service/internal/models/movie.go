package models

import (
	"log"
	"strconv"
	"strings"
)

type Movie struct {
	Id             string
	PrimaryTitle   string
	OriginalTitle  string
	IsAdult        bool
	StartYear      int
	EndYear        int
	RuntimeMinutes int
	Genres         []string
}

func NewMovie(record []string) *Movie {
	var movie Movie

	for i, v := range record {
		if v == "\\N" {
			continue
		}

		switch i {
		case 0:
			movie.Id = v
		case 2:
			movie.PrimaryTitle = v
		case 3:
			movie.OriginalTitle = v
		case 4:
			if v == "0" {
				movie.IsAdult = false
			} else {
				movie.IsAdult = true
			}
		case 5, 6, 7:
			iv, err := strconv.Atoi(v)
			if err != nil {
				log.Println("failed to convert to int value for:", record[0])
				continue
			}

			if i == 5 {
				movie.StartYear = iv
			} else if i == 6 {
				movie.EndYear = iv
			} else {
				movie.RuntimeMinutes = iv
			}
		case 8:
			split := strings.Split(v, ",")
			movie.Genres = append(movie.Genres, split...)
		}
	}

	return &movie
}
