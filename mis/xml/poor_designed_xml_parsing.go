package main

import "fmt"
import "log"
import "encoding/xml"

type Valsi struct {
	Word  string   `xml:"word,attr"`
	Type  string   `xml:"type,attr"`
	Def   string   `xml:"definition"`
	Notes string   `xml:"notes"`
	Class string   `xml:"selmaho"`
	Rafsi []string `xml:"rafsi"`
}

type NLWord struct {
	Word  string `xml:"word,attr"`
	Sense string `xml:"sense,attr"`
	Valsi string `xml:"valsi,attr"`
}

//Whoever made this XML structure needs to be painfully taught a lesson...
type Collection struct {
	From   string   `xml:"from,attr"`
	To     string   `xml:"to,attr"`
	Valsi  []Valsi  `xml:"valsi"`
	NLWord []NLWord `xml:"nlword"`
}

type Vlaste struct {
	Direction []Collection `xml:"direction"`
}

func parseXML(data []byte) Vlaste {
	vlaste := Vlaste{}
	err := xml.Unmarshal(data, &vlaste)
	if err != nil {
		fmt.Println("Problem Decoding!")
		log.Fatal(err)
	}
	return vlaste
}

func main() {
	fmt.Println("Hello, playground")

	input := `
<?xml version="1.0" encoding="UTF-8"?>
<dictionary>

    <direction from="lojban" to="English">
        <valsi word="cipni" type="gismo">
            <rafsi>cpi</rafsi>
            <definition>x1 is a bird of species x2</definition>
            <notes></notes>
        </valsi>
        ...
    </direction>

    <direction from="English" to="lojban">
        <nlword word="eagle" valsi="atkuila" />
        <nlword word="hawk" sense="bird" valsi="aksiptrina" />
        ...
    </direction>

</dictionary>`

	fmt.Printf("%#v\n", parseXML([]byte(input)))
}
