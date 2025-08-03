package opcorpora

import (
	"bytes"
	"encoding/gob"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Dictionary struct {
	XMLName   xml.Name `xml:"dictionary"`
	Text      string   `xml:",chardata"`
	Version   string   `xml:"version,attr"`
	Revision  string   `xml:"revision,attr"`
	Grammemes struct {
		Text     string `xml:",chardata"`
		Grammeme []struct {
			Text        string `xml:",chardata"`
			Parent      string `xml:"parent,attr"`
			Name        string `xml:"name"`
			Alias       string `xml:"alias"`
			Description string `xml:"description"`
		} `xml:"grammeme"`
	} `xml:"grammemes"`
	Restrictions struct {
		Text  string `xml:",chardata"`
		Restr []struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Auto string `xml:"auto,attr"`
			Left struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"left"`
			Right struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"right"`
		} `xml:"restr"`
	} `xml:"restrictions"`
	Lemmata struct {
		Text  string `xml:",chardata"`
		Lemma []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Rev  string `xml:"rev,attr"`
			L    struct {
				Text string `xml:",chardata"`
				T    string `xml:"t,attr"`
				G    []struct {
					Text string `xml:",chardata"`
					V    string `xml:"v,attr"`
				} `xml:"g"`
			} `xml:"l"`
			F []struct {
				Text string `xml:",chardata"`
				T    string `xml:"t,attr"`
				G    []struct {
					Text string `xml:",chardata"`
					V    string `xml:"v,attr"`
				} `xml:"g"`
			} `xml:"f"`
		} `xml:"lemma"`
	} `xml:"lemmata"`
	LinkTypes struct {
		Text string `xml:",chardata"`
		Type []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
		} `xml:"type"`
	} `xml:"link_types"`
	Links struct {
		Text string `xml:",chardata"`
		Link []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			From string `xml:"from,attr"`
			To   string `xml:"to,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
	} `xml:"links"`
}

func LoadDictionaryXML() Dictionary {
	b, err := os.ReadFile("dict.opcorpora.xml") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	var res Dictionary
	err1 := xml.Unmarshal(b, &res)
	if err1 != nil {
		panic(err1)
	}
	return res
}

func StoreDictionary(d Dictionary) {
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	// Encode (send) the value.
	err := enc.Encode(d)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	err1 := os.WriteFile("dict.opcorpora.xml.gob", network.Bytes(), 0777) // just pass the file name
	if err1 != nil {
		fmt.Print(err1)
	}
}

func LoadDictionary() *Dictionary {
	b, err := os.ReadFile("dict.opcorpora.xml.gob") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	bb := bytes.NewBuffer(b)
	dec := gob.NewDecoder(bb) // Will read from network.
	// Decode (receive) the value.
	var q Dictionary
	err1 := dec.Decode(&q)
	if err1 != nil {
		log.Fatal("decode error:", err1)
	}
	return &q
}
