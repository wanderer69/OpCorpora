package opcorpora

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Corpora struct {
	WordDict map[string][]int // ключ - любая(!!!) словоформа но значение - номер записи. // Для поиска словоформы используется внутренний словарь
	WordProp []WordProperties
	Inited   bool
	Debug    int
	Restrict int
}

type WordFormProperties struct {
	ID         string // уникальный идентификатор
	IsBase     bool   // признак базовой формы
	WordForm   string // форма слова
	Properties []string
}

type WordProperties struct {
	Number                 int
	ID                     string                          // уникальный идентификатор
	Rev                    string                          // ревизия
	BaseWordForm           string                          // базовая форма
	WordFormPropertiesDict map[string][]WordFormProperties // список(словарь) свойств форм
}

func StoreCorpora(d *Corpora) {
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	// Encode (send) the value.
	err := enc.Encode(*d)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	err1 := os.WriteFile("opcorpora.gob", network.Bytes(), 0777) // just pass the file name
	if err1 != nil {
		fmt.Print(err1)
	}
}

func LoadCorpora(file_name string) *Corpora {
	b, err := os.ReadFile(file_name) // "opcorpora.gob" just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	bb := bytes.NewBuffer(b)
	dec := gob.NewDecoder(bb) // Will read from network.
	// Decode (receive) the value.
	var q Corpora
	err1 := dec.Decode(&q)
	if err1 != nil {
		log.Fatal("decode error:", err1)
	}
	return &q
}

func ConvertOpenCorpora2HashTable(opd *Dictionary, restricts []string, path string, word string) (*Corpora, bool) {
	co := Corpora{}
	co.WordDict = make(map[string][]int)

	num_cnt := 0
	for _, v := range opd.Lemmata.Lemma {
		var flag bool = true
		var properties WordProperties
		properties.WordFormPropertiesDict = make(map[string][]WordFormProperties)
		properties.Number = num_cnt
		properties.ID = v.ID
		properties.Rev = v.Rev
		properties.BaseWordForm = v.L.T
		wfp := WordFormProperties{}
		wfp.WordForm = v.L.T
		wfp.ID = v.ID
		wfp.IsBase = true
		for _, vv := range v.L.G {
			// ограничения.
			for _, rv := range restricts {
				if vv.V == rv {
					flag = false
				}
			}
			wfp.Properties = append(wfp.Properties, vv.V)
		}
		properties.WordFormPropertiesDict[properties.BaseWordForm] = []WordFormProperties{wfp}

		for _, vvf := range v.F {
			wfp := WordFormProperties{}
			wfp.WordForm = vvf.T
			wfp.ID = v.ID
			wfp.IsBase = false
			for _, vv := range vvf.G {
				wfp.Properties = append(wfp.Properties, vv.V)
			}
			vl, ok := properties.WordFormPropertiesDict[wfp.WordForm]
			if ok {
				properties.WordFormPropertiesDict[wfp.WordForm] = append(vl, wfp)
			} else {
				properties.WordFormPropertiesDict[wfp.WordForm] = []WordFormProperties{wfp}
			}
		}
		if flag {
			co.WordProp = append(co.WordProp, properties)
			pos := len(co.WordProp) - 1
			ll, ok := co.WordDict[properties.BaseWordForm]
			if ok {
				co.WordDict[properties.BaseWordForm] = append(ll, pos)
			} else {
				co.WordDict[properties.BaseWordForm] = []int{pos}
			}
			for _, vvf := range v.F {
				ll, ok := co.WordDict[vvf.T]
				if ok {
					co.WordDict[vvf.T] = append(ll, pos)
				} else {
					co.WordDict[vvf.T] = []int{pos}
				}
			}
			num_cnt = num_cnt + 1
		}
	}
	var ret bool = false
	if num_cnt > 0 {
		ret = true
	}
	return &co, ret
}

func ConvertOpenCorpora2HashTableNew(opd *Dictionary, restricts []string /*, path string, word string*/) (*Corpora, bool) {
	co := Corpora{}
	co.WordDict = make(map[string][]int)

	num_cnt := 0
	for _, v := range opd.Lemmata.Lemma {
		var flag bool = true
		var properties WordProperties
		properties.WordFormPropertiesDict = make(map[string][]WordFormProperties)
		properties.Number = num_cnt
		properties.ID = v.ID
		properties.Rev = v.Rev
		properties.BaseWordForm = v.L.T

		all_props := []string{}
		for _, vv := range v.L.G {
			// ограничения.
			for _, rv := range restricts {
				if vv.V == rv {
					flag = false
				}
			}
			all_props = append(all_props, vv.V)
		}

		for _, vvf := range v.F {
			wfp := WordFormProperties{}
			wfp.WordForm = vvf.T
			wfp.ID = v.ID
			wfp.IsBase = false
			for _, vv := range vvf.G {
				wfp.Properties = append(wfp.Properties, vv.V)
			}
			ap := all_props
			wfp.Properties = append(wfp.Properties, ap...)

			vl, ok := properties.WordFormPropertiesDict[wfp.WordForm]
			if ok {
				properties.WordFormPropertiesDict[wfp.WordForm] = append(vl, wfp)
			} else {
				properties.WordFormPropertiesDict[wfp.WordForm] = []WordFormProperties{wfp}
			}
		}
		if flag {
			co.WordProp = append(co.WordProp, properties)
			num_cnt = num_cnt + 1
		}
	}
	var ret bool = false
	if num_cnt > 0 {
		ret = true
	}
	return &co, ret
}
