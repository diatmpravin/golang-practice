package main

import (
	"encoding/xml"
	"fmt"
)

type Comment struct {
	Paragraph1 string `xml:"p"`
	// Paragraph2 string `xml:"p"`
	// Paragraph3 string `xml:"p"`
}

type Audio struct {
	Status string `xml:"audio>status"`
	PlayedTime string `xml:"audio>played_time"`
	Comments []Comment `xml:"audio>comments"`
}

type Result struct {
	XMLName xml.Name  `xml:"nexgen_audio_export"`
	Audio   `xml:"audio"`
	// Type   string `xml:"audio>type"`
}

func main() {
	xmlResponse := `<?xml version="1.0" encoding="UTF-8"?>
  <nexgen_audio_export>
    <audio ID="id_1667331726_30393658">
      <type>Song</type>
      <status>Playing</status>
      <played_time>09:41:18</played_time>
      <composer>Frederic Delius</composer>
      <title>Violin Sonata No.1</title>
      <artist>Tasmin Little, violin; Piers Lane, piano</artist>
      <comments>
         <p>Comment line1</p>
         <p>Comment <b>line2</b></p>
         <p>Comment line3</p>
      </comments>
    </audio>
  </nexgen_audio_export>`

	v := Result{}
	err := xml.Unmarshal([]byte(xmlResponse), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("XMLName: %#v\n", v)
}
