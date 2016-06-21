package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

var data = `<Data>
	<Series>
		<id>83462</id>
		<Airs_DayOfWeek>Monday</Airs_DayOfWeek>
		<Airs_Time>10:00 PM</Airs_Time>
	</Series>
	<Episode>
		<id>398671</id>
		<Combined_season>1</Combined_season>
		<Director>Rob Bowman</Director>
	</Episode>
</Data>`

type Series struct {
	Id             int    `xml:"Series>id"`
	Airs_DayOfWeek string `xml:"Series>Airs_DayOfWeek"`
	Airs_Time      string `xml:"Series>Airs_Time"`
}

type Episode struct {
	Id              int    `xml:"Episode>id"`
	Combined_season int    `xml:"Episode>Combined_season"`
	Director        string `xml:"Episode>Director"`
}

type Data struct {
	XMLName  xml.Name `xml:"data"`
	Series   `xml:"Series"`
	Episodes []Episode 
}

func main() {
	v := &Data{}
	v.Series = Series{1, "Monday", "10:00 PM"}
	v.Episodes = []Episode{{398671, 1, "Rob Bowman"}, {398672, 2, "Ankit Mishra"}}

	output, err := xml.MarshalIndent(v, "  ", "    ")
	output1, err := xml.Marshal(v)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
	fmt.Println("\n \n")
	os.Stdout.Write(output1)
}
