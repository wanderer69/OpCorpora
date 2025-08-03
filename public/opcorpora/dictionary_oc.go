package opcorpora

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"os"
)

type FormOC struct {
	ID         string   `json:"id"`
	Form       string   `json:"form"`
	Properties []string `json:"propertys"`
}

type LemmaOC struct {
	ID         string            `json:"id"`
	Rev        string            `json:"rev"`
	Base       string            `json:"base"`
	Properties []string          `json:"propertys"`
	Forms      map[string]FormOC `json:"forms"`
}

type LinkOC struct {
	ID   string `json:"id"`
	From string `json:"from"`
	To   string `json:"to"`
	Type string `json:"type"`
}

type TypeOC struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type DictionaryOC struct {
	Lemmas   map[string]LemmaOC // dict ID->Lemma
	BaseToID map[string]string  // dict Base->ID
	Links    map[string]LinkOC  // dict ID->Link
	From2ID  map[string]string  // dict From->ID
	To2ID    map[string]string  // dict To->ID
	Types    map[string]TypeOC  // dict ID->Type
}

func ConvertOpenCorpora2DictionaryOC(opd *Dictionary, restricts []string) (*DictionaryOC, bool) {
	doc := DictionaryOC{}
	doc.Lemmas = make(map[string]LemmaOC)
	doc.BaseToID = make(map[string]string)
	doc.Links = make(map[string]LinkOC)
	doc.From2ID = make(map[string]string)
	doc.To2ID = make(map[string]string)
	doc.Types = make(map[string]TypeOC)
	for _, v := range opd.Lemmata.Lemma {
		loc := LemmaOC{}
		loc.Forms = make(map[string]FormOC)
		loc.ID = v.ID
		loc.Rev = v.Rev
		loc.Base = v.L.T

		flag := true
		for _, vv := range v.L.G {
			// ограничения.
			for _, rv := range restricts {
				if vv.V == rv {
					flag = false
				}
			}
			if flag {
				loc.Properties = append(loc.Properties, vv.V)
			}
		}
		if flag {
			for _, vvf := range v.F {
				foc := FormOC{}
				foc.Form = vvf.T
				foc.ID = v.ID
				for _, vv := range vvf.G {
					foc.Properties = append(foc.Properties, vv.V)
				}
				loc.Forms[foc.Form] = foc
			}
			doc.Lemmas[loc.ID] = loc
			doc.BaseToID[loc.Base] = loc.ID
		}
	}
	for _, v := range opd.LinkTypes.Type {
		toc := TypeOC{v.ID, v.Text}
		doc.Types[v.ID] = toc
	}
	for _, v := range opd.Links.Link {
		loc := LinkOC{}
		loc.ID = v.ID
		loc.From = v.From
		loc.To = v.To
		loc.Type = v.Type
		flag := true
		_, ok1 := doc.Lemmas[loc.From]
		if !ok1 {
			flag = false
		}
		_, ok2 := doc.Lemmas[loc.To]
		if !ok2 {
			flag = false
		}
		_, ok3 := doc.Types[loc.Type]
		if !ok3 {
			flag = false
		}
		if flag {
			doc.Links[loc.ID] = loc
			doc.From2ID[loc.From] = loc.ID
			doc.To2ID[loc.To] = loc.ID
		}
	}
	return &doc, true
}

func StoreDictionaryOC(d *DictionaryOC) {
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	// Encode (send) the value.
	err := enc.Encode(*d)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	err1 := os.WriteFile("dict.opcorpora.json.gob", network.Bytes(), 0777) // just pass the file name
	if err1 != nil {
		fmt.Print(err1)
	}
}

func LoadDictionaryOC() *DictionaryOC {
	b, err := os.ReadFile("dict.opcorpora.json.gob") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	bb := bytes.NewBuffer(b)
	dec := gob.NewDecoder(bb) // Will read from network.
	// Decode (receive) the value.
	var q DictionaryOC
	err1 := dec.Decode(&q)
	if err1 != nil {
		log.Fatal("decode error:", err1)
	}
	return &q
}

func FindDictionaryOC(d *DictionaryOC, link_type string) ([]string, error) {
	res := []string{}
	for _, v := range d.Links {
		if v.Type == link_type {
			lf, ok1 := d.Lemmas[v.From]
			if !ok1 {
				return res, errors.New("from not found lemmas")
			}
			lt, ok2 := d.Lemmas[v.To]
			if !ok2 {
				return res, errors.New("to not found lemmas")
			}
			ss := fmt.Sprintf("lf %v\t lt %v", lf.Base, lt.Base)
			res = append(res, ss)
		}
	}
	return res, nil
}

func FindLinkToFromDictionaryOC(d *DictionaryOC, link_type string, to string) (string, error) {
	res := ""
	id, ok := d.To2ID[to]
	if ok {
		lf, ok1 := d.Lemmas[id]
		if ok1 {
			return lf.Base, errors.New("from not found lemmas")
		}
	}
	return res, nil
}

// поиск базовой формы для имени существительного и имени
func FindBaseFormToFromDictionaryOC(d *DictionaryOC, link_type string, to string) (string, error) {
	res := ""
	id, ok := d.To2ID[to]
	if ok {
		lf, ok1 := d.Lemmas[id]
		if ok1 {
			return lf.Base, errors.New("from not found lemmas")
		}
	}
	return res, nil
}
