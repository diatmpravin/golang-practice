package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var data = `{
	"Series": {
		"id": 1,
		"Airs_DayOfWeek": "Monday",
		"Airs_Time": "10:00 PM"
	},
	"Episodes": [{
		"id": 398671,
		"Combined_season": 1,
		"Director": "Rob Bowman"
	}, {
		"id": 398672,
		"Combined_season": 2,
		"Director": "Ankit Mishra"
	}]
}`

type Series struct {
	Id             int    `json:"id"`
	Airs_DayOfWeek string `json:"Airs_DayOfWeek"`
	Airs_Time      string `json:"Airs_Time"`
}

type Episode struct {
	Id              int    `json:"id"`
	Combined_season int    `json:"Combined_season"`
	Director        string `json:"Director"`
}

type Data struct {
	Series   `json:"Series"`
	Episodes []Episode
}

func main() {
	v := &Data{}
	v.Series = Series{1, "Monday", "10:00 PM"}
	v.Episodes = []Episode{{398671, 1, "Rob Bowman"}, {398672, 2, "Ankit Mishra"}}

	output, err := json.MarshalIndent(v, "  ", "    ")
	output1, err := json.Marshal(v)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
	fmt.Println("\n \n")
	os.Stdout.Write(output1)
}
