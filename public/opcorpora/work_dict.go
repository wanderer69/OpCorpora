package opcorpora

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"

	"errors"
	"fmt"

	//"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	//. "github.com/wanderer69/SmallDB/v3"
	"github.com/wanderer69/SmallDB/public/index"
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
	//WordForm        string
	//Properties      []string
	//PropertiesDict  map[string]int // список свойств общих для всех словоформ
}

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
	// fmt.Printf("%+v\n", res)
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

func LoadDictionary() Dictionary {
	b, err := os.ReadFile("dict.opcorpora.xml.gob") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	bb := bytes.NewBuffer(b)
	// var network bytes.Buffer        // Stand-in for a network connection
	dec := gob.NewDecoder(bb) // Will read from network.
	// Decode (receive) the value.
	var q Dictionary
	err1 := dec.Decode(&q)
	if err1 != nil {
		log.Fatal("decode error:", err1)
	}
	return q
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
	// var network bytes.Buffer        // Stand-in for a network connection
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
		//d.Links[k].ID
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
	// var network bytes.Buffer        // Stand-in for a network connection
	dec := gob.NewDecoder(bb) // Will read from network.
	// Decode (receive) the value.
	var q Corpora
	err1 := dec.Decode(&q)
	if err1 != nil {
		log.Fatal("decode error:", err1)
	}
	return &q
}

type Settings struct {
	RestrictionsList []string `json:"restrictions_list"`
}

func InitSettings() *Settings {
	settings := Settings{}
	settings.RestrictionsList = []string{"Abbr", "Name", "Surn", "Patr", "Geox", "Orgn", "Trad", "Init", "Hypo"}
	return &settings
}

func LoadSettings(file_name string) (*Settings, error) {
	data, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	// json data
	var settings Settings
	// unmarshall it
	err = json.Unmarshal(data, &settings)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	return &settings, nil
}

func SaveSettings(file_name string, settings *Settings) error {
	// unmarshall it
	data, err := json.Marshal(&settings)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	_ = os.WriteFile(file_name, data, 0644)
	return nil
}

func LoadBadWordsList() ([]string, error) {
	rsl := []string{}
	b, err := os.ReadFile("bad_words.txt")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	sl := strings.Split(string(b), "\n")
	for i := range sl {
		s := strings.Trim(sl[i], " \r\n\t")
		rsl = append(rsl, s)
	}
	return rsl, nil
}

func ConvertOpenCorpora2HashTable(opd *Dictionary, restricts []string, path string, word string) (*Corpora, bool) {
	co := Corpora{}
	co.WordDict = make(map[string][]int)

	num_cnt := 0
	for _, v := range opd.Lemmata.Lemma {
		//fmt.Printf("%+v\r\n", v)
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

		/*
		   if word == v.L.T {
		       fmt.Printf("v.L %v v.F %v\r\n", v.L, v.F)
		   }
		*/

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

			/*
			   if word == vvf.T {
			       fmt.Printf("vvf %v\r\n", vvf)
			   }
			*/

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
			/*
			   if word == v.L.T {
			       fmt.Printf("pos %v properties %v\r\n", pos, properties)
			   }
			*/
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
			/*
			   if word == v.L.T {
			       fmt.Printf("co.WordProp %v\r\n", co.WordProp[pos])
			   }
			*/
			num_cnt = num_cnt + 1
		}
		/*
		   if word == v.L.T {
		       ll, ok := co.WordDict[word]
		       if ok {
		           fmt.Printf("ll -> %v\r\n", ll)
		       } else {
		       }
		   }
		*/
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
		//fmt.Printf("%+v\r\n", v)
		var flag bool = true
		var properties WordProperties
		properties.WordFormPropertiesDict = make(map[string][]WordFormProperties)
		properties.Number = num_cnt
		properties.ID = v.ID
		properties.Rev = v.Rev
		properties.BaseWordForm = v.L.T
		/*
			wfp := WordFormProperties{}
			wfp.WordForm = v.L.T
			wfp.ID = v.ID
			wfp.IsBase = true
		*/
		/*
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
		   	//WordForm        string
		   	//Properties      []string
		   	//PropertiesDict  map[string]int // список свойств общих для всех словоформ
		   }
		*/

		all_props := []string{}
		for _, vv := range v.L.G {
			// ограничения.
			for _, rv := range restricts {
				if vv.V == rv {
					flag = false
				}
			}
			//wfp.Properties = append(wfp.Properties, vv.V)
			all_props = append(all_props, vv.V)
		}
		/*
			bwfp := []WordFormProperties{wfp}
			IsBase     bool   // признак базовой формы
			WordForm   string // форма слова
			Properties []string

			properties.WordFormPropertiesDict[properties.BaseWordForm] =
		*/

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
			/*
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
			*/
			num_cnt = num_cnt + 1
		}
		/*
		   if word == v.L.T {
		       ll, ok := co.WordDict[word]
		       if ok {
		           fmt.Printf("ll -> %v\r\n", ll)
		       } else {
		       }
		   }
		*/
	}
	var ret bool = false
	if num_cnt > 0 {
		ret = true
	}
	return &co, ret
}

type OCorpora struct {
	Sdb   *index.SmallDB
	Debug int
}

func OpenOCorpora(path string) (*OCorpora, error) {
	sdb := index.InitSmallDB(path) // "./OCorporaDB"
	sdb.Debug = 0
	if !sdb.Inited {
		fl := []string{"word_form", "word", "word_property"}
		err := sdb.CreateDB(fl, path) //  "./OCorporaDB"
		if err != nil {
			fmt.Printf("Error creating DB %v\r\n", err)
			return nil, err
		}
		sdb.CreateIndex([]string{"word_form"})
		sdb.CreateIndex([]string{"word"})
		fmt.Printf("CreateDB end\r\n")
		sdb = index.InitSmallDB(path) // "./OCorporaDB"

	}
	sdb.OpenDB()
	oc := &OCorpora{}
	oc.Sdb = sdb
	return oc, nil
}

func OpenOCorporaFull(path string, debug int) (*OCorpora, error) {
	sdb := index.InitSmallDB(path) // "./OCorporaDBFull"
	sdb.Debug = 0
	Global_Dict_Init()
	if debug > 2 {
		fmt.Printf("Inited %v\r\n", sdb.Inited)
	}
	if !sdb.Inited {
		fl := []string{"word_form", "word", "POS", "animacy", "aspects", "cases", "genders", "involvement", "moods", "numbers", "persons", "tenses", "transitivity", "voices", "word_property"}
		err := sdb.CreateDB(fl, path) //  "./OCorporaDB"
		if err != nil {
			fmt.Printf("Error creating DB %v\r\n", err)
			return nil, err
		}
		sdb.CreateIndex([]string{"word_form"})
		sdb.CreateIndex([]string{"word"})
		fmt.Printf("CreateDB end\r\n")
		sdb = index.InitSmallDB(path) // "./OCorporaDBFull"
	}
	sdb.OpenDB()
	oc := &OCorpora{}
	oc.Sdb = sdb
	return oc, nil
}

func OpenOCorporaShort(path_db string) (*OCorpora, error) {
	sdb := index.InitSmallDB(path_db) // "./OCorporaDBShort"
	sdb.Debug = 0
	if !sdb.Inited {
		fl := []string{"word_form", "word", "word_property"} // слово-форма, базовая форма, локальные свойства
		err := sdb.CreateDB(fl, path_db)                     //  "./OCorporaDB"
		if err != nil {
			fmt.Printf("Error creating DB %v\r\n", err)
			return nil, err
		}
		sdb.CreateIndex([]string{"word_form"})
		sdb.CreateIndex([]string{"word"})
		fmt.Printf("CreateDB end\r\n")
		sdb = index.InitSmallDB(path_db) // "./OCorporaDBShort"
	}
	sdb.OpenDB()
	oc := &OCorpora{}
	oc.Sdb = sdb
	return oc, nil
}

func OpenYo(path string) (*OCorpora, error) {
	sdb := index.InitSmallDB(path)
	sdb.Debug = 0
	Global_Dict_Init()
	if !sdb.Inited {
		fl := []string{"word_wo_yo", "word_w_yo"}
		err := sdb.CreateDB(fl, path) //  "./OCorporaDB"
		if err != nil {
			fmt.Printf("Error creating DB %v\r\n", err)
			return nil, err
		}
		sdb.CreateIndex([]string{"word_wo_yo"})
		sdb.CreateIndex([]string{"word_w_yo"})
		fmt.Printf("CreateDB end\r\n")
		sdb = index.InitSmallDB(path) // "./OCorporaDBFull"
	}
	sdb.OpenDB()
	oc := &OCorpora{}
	oc.Sdb = sdb
	return oc, nil
}

func (oc *OCorpora) StoreToSmallDB(co *Corpora) {
	for i := range co.WordProp {
		wp := co.WordProp[i]
		// добавляем в базу данных
		mCache := new(bytes.Buffer)
		encCache := gob.NewEncoder(mCache)
		encCache.Encode(wp)
		wpb := mCache.String()
		word_data := wp.BaseWordForm
		for k := range wp.WordFormPropertiesDict {
			word_form_data := k
			oc.Sdb.StoreRecord(word_form_data, word_data, " ")
		}
		oc.Sdb.StoreRecord(" ", word_data, wpb)
		if i == co.Restrict {
			break
		}
	}
}

func (oc *OCorpora) StoreToSmallDBFull(co *Corpora, bad_words_list []string) {
	bwd := map[string]bool{}
	for i := range bad_words_list {
		s := bad_words_list[i]
		bwd[s] = true
	}
	fmt.Printf("%v\r\n", bwd)
	data_a := []map[string]string{}
	for i := range co.WordProp {
		wp := co.WordProp[i]
		wpd := wp.WordFormPropertiesDict

		// добавляем в базу данных
		mCache := new(bytes.Buffer)
		encCache := gob.NewEncoder(mCache)
		encCache.Encode(wp)
		//wpb := string(mCache.Bytes())
		word_data := wp.BaseWordForm
		//data := make(map[string]string)
		for k := range wp.WordFormPropertiesDict {
			word_form_data := k

			_, flag_bw := bwd[word_form_data]

			wpi := wpd[word_form_data]
			//fmt.Printf("word_data %v, word_form_data %v, flag_bw %v, wpi %v\r\n", word_data, word_form_data, flag_bw, wpi)
			for j := range wpi {
				args := make(map[string]string)
				args["word_form"] = word_form_data
				args["word"] = word_data
				props := []string{}
				flag_form := true
				for m := range wpi[j].Properties {
					attr, tag := Tag2str_attr_int(wpi[j].Properties[m])
					if attr == "" {
						fmt.Printf("Error! in %v %v %v not attr (%v)\r\n", word_form_data, word_data, wpi[j].Properties[m], wpi[j].Properties)
						// return
					} else {
						if attr == "POS" {
							flag_form = false
						}
						if attr != "attributes" {
							args[attr] = tag
						}
						//						data[attr] = tag
					}
					props = append(props, wpi[j].Properties[m])
				}
				wpi_n := wpd[word_data]
				pos := ""
				if flag_form {
					for l := range wpi_n {
						for n := range wpi_n[l].Properties {
							attr, tag := Tag2str_attr_int(wpi_n[l].Properties[n])
							if attr == "POS" {
								args[attr] = tag
								pos = wpi_n[l].Properties[n]
							}
						}
					}
					props = append(props, pos)
				}
				prop, _ := json.MarshalIndent(props, "", "  ") // wpi[j].Properties

				args["word_property"] = string(prop)

				//				data["word_form"] = word_form_data
				//				data["word"] = word_data
				//				data["word_property"] = string(prop)

				res1, num1, err1 := oc.Sdb.StoreRecordOnMap(args)
				if err1 != nil {
					fmt.Printf("err store %v\r\n", err1)
					return
				}
				if res1 < 0 {
					fmt.Printf("err store 2 num %v res %v\r\n", num1, res1)
					// return
					continue
				}

				if false {
					if flag_bw {
						fmt.Printf("word_form_data %v word_data %v wpi %#v\r\n", word_form_data, word_data, wpi[j].Properties)
						fmt.Printf("args %v\r\n", args)

						res1, num1, err1 := oc.Sdb.StoreRecordOnMap(args)
						if err1 != nil {
							fmt.Printf("err store %v\r\n", err1)
							return
						}
						if res1 < 0 {
							fmt.Printf("err store 2 num %v res %v\r\n", num1, res1)
							return
							//continue
						}
						fmt.Printf("res1 %v\r\n", res1)

						wfdl := []string{word_form_data}
						_, num, err := oc.Sdb.FindRecordIndexString([]string{"word_form"}, wfdl)
						// fmt.Printf("ds %v num %v, err %v\r\n", ds, num, err)
						if err != nil {
							fmt.Printf("Error 1 %v num %v %v\r\n", err, num, wfdl)
							return
							//continue
						}
						if num != 0 {
							fmt.Printf("Error 2 %v %v\r\n", num, wfdl)
							return
							//continue
						}
						/*
							if len(ds) < len(wpi) {

							}
						*/
					}
				}
				//				data_a = append(data_a, data)
				//oc.Sdb.Store_record(word_form_data, word_data, " ")
			}
		}
		// fmt.Printf("word_data %v wpb %#v\r\n", word_data, wp)
		// fmt.Printf("word_data %v\r\n", word_data)
		// oc.Sdb.Store_record(" ", word_data, wpb)
		if i == co.Restrict {
			break
		}
	}
	prop, _ := json.MarshalIndent(data_a, "", "  ")
	err1 := os.WriteFile("writed.json", prop, 0777)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
}

func (oc *OCorpora) ConvertToYo(co *Corpora) {
	data_a := make(map[string]string)
	check_yo := func(text string) (string, bool) {
		sn := ""
		flag := false
		for i := 0; i < len(text); {
			runeValue, width := utf8.DecodeRuneInString(text[i:])
			if false {
				fmt.Printf("%#U starts at byte position %d %v\n", runeValue, i, string(runeValue))
			}
			ss := string(runeValue)
			if ss == "ё" {
				sn = sn + "е"
				flag = true
			} else {
				sn = sn + ss
			}
			i = i + width
		}
		return sn, flag
	}
	store := func(wo_word string, w_word string) {
		args := make(map[string]string)
		args["word_wo_yo"] = wo_word
		args["word_w_yo"] = w_word
		data_a[wo_word] = w_word
		res1, num1, err1 := oc.Sdb.StoreRecordOnMap(args)
		if err1 != nil {
			fmt.Printf("err store %v\r\n", err1)
			return
		}
		if res1 < 0 {
			fmt.Printf("err store 2 num %v res %v\r\n", num1, res1)
		}
	}
	for i := range co.WordProp {
		wp := co.WordProp[i]
		mCache := new(bytes.Buffer)
		encCache := gob.NewEncoder(mCache)
		encCache.Encode(wp)
		word_data := wp.BaseWordForm
		wo_yo, ok := check_yo(word_data)
		if ok {
			store(wo_yo, word_data)
		}
		for k := range wp.WordFormPropertiesDict {
			word_form_data := k
			wo_yo, ok := check_yo(word_form_data)
			if ok {
				store(wo_yo, word_form_data)
			}
		}
		if i == co.Restrict {
			break
		}
	}
	prop, _ := json.MarshalIndent(data_a, "", "  ")
	err1 := os.WriteFile("writed.json", prop, 0777)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
}

func Init_Unique_Value() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func Unique_Value(len_n int) string {
	var bytes_array []byte

	for i := 0; i < len_n; i++ {
		bytes := rand.Intn(35)
		if bytes > 9 {
			bytes = bytes + 7
		}
		bytes_array = append(bytes_array, byte(bytes+16*3))
	}
	str := string(bytes_array)
	return str
}

func (oc *OCorpora) StoreToSmallDBShort(co *Corpora) {
	//cnt := 30
	Init_Unique_Value()
	for i := range co.WordProp {
		/*
			if i == cnt {
				//            return
			}
		*/
		wp := co.WordProp[i]
		// добавляем в базу данных
		/*
		   mCache := new(bytes.Buffer)
		   encCache := gob.NewEncoder(mCache)
		   encCache.Encode(wp)
		   wpb := string(mCache.Bytes())
		*/
		word_data := wp.BaseWordForm
		wp_data, err := json.Marshal(&wp)
		if err != nil {
			fmt.Println("error:", err)
			//return err
		}
		oc.Sdb.StoreRecord("fake_"+Unique_Value(7), word_data, string(wp_data))
		for k, v := range wp.WordFormPropertiesDict {
			word_form_data := k
			local_data, err := json.Marshal(&v)
			if err != nil {
				fmt.Println("error:", err)
				//return err
			}
			oc.Sdb.StoreRecord(word_form_data, word_data+"_"+Unique_Value(7), string(local_data))
		}
		// проверяем
		wfdl := []string{word_data}
		ds, rec_id, err1 := oc.Sdb.FindRecordIndexString([]string{"word"}, wfdl)
		if err1 != nil {
			fmt.Printf("word_data %v\r\n", word_data)
			fmt.Printf("ds %v num %v err %v\r\n", ds, rec_id, err1)
			fmt.Printf("Error word!! %v\r\n", word_data)
		} else {
			for k := range wp.WordFormPropertiesDict {
				wfdl := []string{k}
				ds2, num, err2 := oc.Sdb.FindRecordIndexString([]string{"word_form"}, wfdl)
				if err2 != nil {
					fmt.Printf("ds2 %v num %v err2 %v\r\n", ds2, num, err2)
					fmt.Printf("Error word!! %v %v\r\n", word_data, k)
				}
			}
		}
		/*
		   res1, res2 := FindWordCorpora2(word_data, oc)
		   if res1 != word_data {
		       fmt.Printf("Error word!! %v %v %v", word_data, res1, res2)
		       return
		   } else {
		       for k := range wp.WordFormPropertiesDict {
		           word_form_data := k
		           res1, res2 := FindWordCorpora1(word_form_data, oc)
		           if res1 != word_data {
		               fmt.Printf("Error form!! %v %v %v %v", word_form_data, res1, word_data, res2)
		               return
		           }
		       }
		   }
		*/
	}
}

func FindWord(word string, co *Corpora) {
	fmt.Printf("word %v\r\n", word)
	ll, ok := co.WordDict[word]
	if ok {
		fmt.Printf("ll %v\r\n", ll)
		for i := range ll {
			wp := co.WordProp[ll[i]]
			fmt.Printf("wp %v\r\n", wp)
		}
	} else {
		fmt.Printf("None\r\n")
	}
}

func FindDictionary(opd *Dictionary, word string) []string {
	res := []string{}
	for _, v := range opd.Lemmata.Lemma {
		//var flag bool = true
		if v.L.T == word {
			for _, vv := range v.L.G {
				res = append(res, vv.V)
				fmt.Printf("v.F %v\r\n", v.F)
				return res
			}
		}
		for _, vvf := range v.F {
			if word == vvf.T {
				for _, vv := range vvf.G {
					fmt.Printf("v.L %v v.F %v\r\n", v.L, v.F)
					res = append(res, vv.V)
				}
				return res
			}
		}
	}
	return res
}

func (oc *OCorpora) TestFindSmallDBFull(co *Corpora) {
	fmt.Printf("TestFindSmallDBFull begin\r\n")
	flag_a := 0
	for i := range co.WordProp {
		wp := co.WordProp[i]
		wpd := wp.WordFormPropertiesDict

		// добавляем в базу данных
		mCache := new(bytes.Buffer)
		encCache := gob.NewEncoder(mCache)
		encCache.Encode(wp)
		//wpb := string(mCache.Bytes())
		word_data := wp.BaseWordForm
		flag_m := 0
		for k := range wp.WordFormPropertiesDict {
			word_form_data := k
			wpi := wpd[word_form_data]

			wfdl := []string{word_form_data}
			ds, num, err := oc.Sdb.FindRecordIndexString([]string{"word_form"}, wfdl)
			// fmt.Printf("ds %v num %v, err %v\r\n", ds, num, err)
			if err != nil {
				fmt.Printf("Error 1 %v num %v %v\r\n", err, num, wfdl)
				//return
				continue
			}
			if num != 0 {
				fmt.Printf("Error 2 %v %v\r\n", num, wfdl)
				//return
				continue
			}

			if len(ds) < len(wpi) {
				fmt.Printf("Warning 3! Length ds %v not equal length wpi %v!\r\n", len(ds), len(wpi))
				for i := range ds {
					dsi := ds[i]
					fmt.Printf("ds %v\r\n\r\n", *dsi)
				}
				for i := range wpi {
					fmt.Printf("wpi[%v] %v\r\n", i, wpi[i])
				}
			}
			flag_w := 0
			for j := range wpi {
				// args := make(map[string]string)
				// args["word_form"] = word_form_data
				// args["word"] = word_data
				flag := 0
				for n_rec := range ds {
					rec := ds[n_rec]
					val, err := oc.Sdb.GetFieldValueByName(rec, "word")
					if err != nil {
						fmt.Printf("Error 4! %v\r\n", err)
						return
					} else {
						if val != word_data {
							//fmt.Printf("Error! val %v not equal tag %v\r\n", val, word_data)
						} else {
							flag_p := 0
							cmp := 0
							for m := range wpi[j].Properties {
								attr, tag := Tag2str_attr_int(wpi[j].Properties[m])
								if len(attr) != 0 {
									//fmt.Printf("attr %v")
									if ((((((((((attr == "POS" && attr == "animacy") && attr == "aspects") && attr == "cases") && attr == "genders") &&
										attr == "involvement") && attr == "moods") && attr == "numbers") && attr == "persons") && attr == "tenses") &&
										attr == "transitivity") && attr != "voices" {
										/*
											if attr != "POS" || attr != "animacy" || attr != "aspects" || attr != "cases" || attr != "genders" ||
												attr != "involvement" || attr != "moods" || attr != "numbers" || attr != "persons" || attr != "tenses" ||
												attr != "transitivity" || attr != "voices" {
											} else {
										*/
										val, err := oc.Sdb.GetFieldValueByName(rec, attr)
										if err != nil {
											fmt.Printf("Error 5! %v attr %v\r\n", err, attr)
											return
										} else {
											if val != tag {
												// fmt.Printf("Error! val %v not equal tag %v\r\n", val, tag)
											} else {
												flag_p = flag_p + 1
											}
										}
									}
									cmp = cmp + 1

								}
							}
							if cmp == flag_p {
								flag = flag + 1
							}
						}
					}
				}
				if flag == 0 {
					//prop, _ := json.MarshalIndent(wpi[j].Properties, "", "  ")
					//args["word_property"] = string(prop)
					//fmt.Printf("word_form_data %v word_data %v wpi %#v\r\n", word_form_data, word_data, wpi[j].Properties)
				} else {
					flag_w = flag_w + 1
				}
			}
			if flag_w == len(wpi) {
				flag_m = flag_m + 1
			}
		}
		if flag_m == len(wp.WordFormPropertiesDict) {
			flag_a = flag_a + 1
		}
		// fmt.Printf("word_data %v wpb %#v\r\n", word_data, wp)
		// fmt.Printf("word_data %v\r\n", word_data)
		// oc.Sdb.Store_record(" ", word_data, wpb)
		if i == co.Restrict {
			break
		}
	}
	fmt.Printf("flag_a %v, len(co.WordProp) %v\r\n", flag_a, len(co.WordProp))
}

func (oc *OCorpora) TestFindSmallDB(co *Corpora) {
	fmt.Printf("TestFindSmallDB begin\r\n")
	for i := range co.WordProp {
		wp := co.WordProp[i]
		// добавляем в базу данных
		mCache := new(bytes.Buffer)
		encCache := gob.NewEncoder(mCache)
		encCache.Encode(wp)
		//wpb := string(mCache.Bytes())
		word_data := wp.BaseWordForm

		// FindWordCorpora(word string, oc)

		for k := range wp.WordFormPropertiesDict {
			word_form_data := k
			fmt.Printf("word form %v word %v\r\n", word_form_data, word_data)
			FindWordCorporaTest(word_form_data, word_data, oc)
			//oc.Sdb.Store_record(word_form_data, word_data, " ")
		}
		//oc.Sdb.Store_record(" ", word_data, wpb)
		//fmt.Printf("word %v, word data %v\r\n", word_data, wpb)
		if i == 1000 {
			break
		}
	}
}

func FindWordCorporaTest(word string, word_data string, oc *OCorpora) {
	wfdl := []string{word}
	ds, num, err := oc.Sdb.FindRecordIndexString([]string{"word_form"}, wfdl)
	//	fmt.Printf("ds %v num %v, err %v\r\n", ds, num, err)
	if err != nil {
		fmt.Printf("Error %v\r\n", err)
		return
	}
	if num == 0 {
		if len(ds) > 0 {
			for k := range ds {
				res := ds[k].FieldsValue
				if len(res) > 0 {
					if res[0] != word {
						fmt.Printf("Error! %v !=  %v\r\n", res[0], word)
						return
					}
					wsl := strings.Split(res[1], "_")
					word_data1 := []string{wsl[0]}
					fmt.Printf("word_data1 %v\r\n", word_data1)
					ds1, num1, err1 := oc.Sdb.FindRecordIndexString([]string{"word"}, word_data1)
					if err1 != nil {
						fmt.Printf("Error record %v %v\r\n", num1, err1)
						return
					}
					// fmt.Printf("ds1 %v num1 %v err1 %v\r\n", ds1, num1, err1)
					fmt.Printf("len(ds1) %v num1 %v\r\n", len(ds1), num1)
					if false {
						if len(ds1) > 0 {
							for i := range ds1 {
								if len(ds1[i].FieldsValue) > 0 {
									res := ds1[i].FieldsValue
									if len(res) > 0 {
										fmt.Printf("word_data %v\r\n", res[1])
										if false {
											wdb := strings.Trim(wsl[0], " \t\r\n")
											var prop WordProperties
											pCache := bytes.NewBuffer([]byte(wdb))
											decCache := gob.NewDecoder(pCache)
											decCache.Decode(&prop)
											fmt.Printf("%v\r\n", prop)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func FindWordCorpora(word string, oc *OCorpora) {
	wfdl := []string{word}
	ds, num, err := oc.Sdb.FindRecordIndexString([]string{"word_form"}, wfdl)
	fmt.Printf("ds %v num %v, err %v\r\n", ds, num, err)
	if err != nil {
		fmt.Printf("Error %v\r\n", err)
		return
	}
	if num == 0 {
		if len(ds) > 0 {
			res := ds[0].FieldsValue
			if len(res) > 0 {
				wsl := strings.Split(res[1], "_")
				word_data1 := []string{wsl[0]}
				ds1, num1, err1 := oc.Sdb.FindRecordIndexString([]string{"word"}, word_data1)
				if err1 != nil {
					fmt.Printf("Error %v\r\n", err1)
					return
				}
				fmt.Printf("ds1 %v num1 %v err1 %v\r\n", ds1, num1, err1)
				if len(ds1) > 0 {
					for i := range ds1 {
						if len(ds1[i].FieldsValue) > 0 {
							res := ds1[i].FieldsValue
							if len(res) > 0 {
								wdb := strings.Trim(wsl[0], " \t\r\n")
								var prop WordProperties
								pCache := bytes.NewBuffer([]byte(wdb))
								decCache := gob.NewDecoder(pCache)
								decCache.Decode(&prop)
								fmt.Printf("%v\r\n", prop)
							}
						}
					}
				}
			}
		}
	}
}

func FindWordCorporaFull(word string, oc *OCorpora) ([]ResultWordProp, error) {
	result := []ResultWordProp{}
	wfdl := []string{word}
	ds, num, err := oc.Sdb.FindRecordIndexString([]string{"word_form"}, wfdl)
	if err != nil {
		if oc.Debug > 2 {
			fmt.Printf("Error %v %v\r\n", num, err)
		}
		return result, err
	}
	if oc.Debug > 4 {
		fmt.Printf("ds %v num %v err %v\r\n", ds, num, err)
	}
	if len(ds) > 0 {
		for i := range ds {
			rec := ds[i]
			//			res := ds[i].FieldsValue
			//			if len(res) > 0 {
			//				if oc.Debug > 4 {
			//					fmt.Printf("res %v\r\n", res)
			//				}

			wordform, err := oc.Sdb.GetFieldValueByName(rec, "word_form")
			if err != nil {
				fmt.Printf("Error! %v\r\n", err)
				return result, err
			}

			word, err := oc.Sdb.GetFieldValueByName(rec, "word")
			if err != nil {
				fmt.Printf("Error! %v\r\n", err)
				return result, err
			}
			word_property, err := oc.Sdb.GetFieldValueByName(rec, "word_property")
			if err != nil {
				fmt.Printf("Error! %v\r\n", err)
				return result, err
			}
			// json data
			var arr []string
			// unmarshall it
			err1 := json.Unmarshal([]byte(word_property), &arr)
			if err1 != nil {
				fmt.Println("error:", err1)
				return result, err1
			}
			rwp := ResultWordProp{}
			rwp.Word = wordform
			rwp.BaseWord = word

			result = append(result, rwp)
			//			}
		}
	}
	return result, nil
}

func FindWordCorporaService(oc *OCorpora, word string) ([][][]string, error) {
	wfdl := []string{word}
	result := [][][]string{}
	ds, num, err := oc.Sdb.FindRecordIndexString([]string{"word_form"}, wfdl)
	//fmt.Printf("ds %v num %v, err %v\r\n", ds, num, err)
	if err != nil {
		fmt.Printf("Error %v\r\n", err)
		return result, err
	}
	if num == 0 {
		if len(ds) > 0 {
			for i := range ds {
				if len(ds[i].FieldsValue) > 0 {
					rec, err1 := oc.Sdb.GetFieldsValueWithName(ds[i])
					if err1 != nil {
						fmt.Printf("Error %v\r\n", err1)
						return result, err1
					}
					result = append(result, rec)
				}
			}
		}
	}
	return result, nil
}

func FindWordCorporaShort(word string, oc *OCorpora) [][]string {
	wfdl := []string{word}
	ds, num, err := oc.Sdb.FindRecordIndexString([]string{"word_form"}, wfdl)
	if err != nil {
		fmt.Printf("Error %v\r\n", err)
		return [][]string{}
	}
	fmt.Printf("ds %v num %v err %v\r\n", ds, num, err)

	result := [][]string{}
	if num == 0 {
		if len(ds) > 0 {
			for i := range ds {
				res := ds[i].FieldsValue
				if len(res) > 0 {
					wsl := strings.Split(res[1], "_")
					word_data1 := []string{wsl[0]}
					ds1, num1, err1 := oc.Sdb.FindRecordIndexString([]string{"word"}, word_data1)
					if err1 != nil {
						fmt.Printf("Error %v\r\n", err1)
						return [][]string{}
					}

					if num1 == 0 {
						fmt.Printf("ds1 %v num1 %v\r\n", ds1, num1)
						if len(ds1) > 0 {
							for j := range ds1 {
								if len(ds1[j].FieldsValue) > 0 {
									res := ds1[j].FieldsValue
									if len(res) > 0 {
										if res[1] == wsl[0] {
											data := res[2]
											// json data
											var prop WordProperties
											// unmarshall it
											err1 := json.Unmarshal([]byte(data), &prop)
											if err1 != nil {
												fmt.Println("error:", err1)
												//return nil, err1
											}
											// fmt.Printf("prop %+v\r\n", prop)
											_, ok := prop.WordFormPropertiesDict[word]
											if ok {
												result = append(result, []string{word, wsl[0], data})
											}
										}
									}
								}
							}
						}
					} else {
						fmt.Printf("ds1 %v err1 %v\r\n", ds1, err1)
					}
				}
			}
		}
	}
	return result
}

func GetWordProperties(data string) WordProperties {
	var prop WordProperties
	// unmarshall it
	err1 := json.Unmarshal([]byte(data), &prop)
	if err1 != nil {
		fmt.Println("error:", err1)
		//return nil, err1
	}
	return prop
}

type ResultWordProp struct {
	Word       string
	BaseWord   string
	Properties []string
}

func GetWordFromWordProperties(word string, prop WordProperties) []ResultWordProp {
	rwp := []ResultWordProp{}
	base_word := prop.BaseWordForm
	rwp_b := ResultWordProp{}
	//	bb := make(map[string]ResultWordProp)
	bw_prop, ok := prop.WordFormPropertiesDict[base_word]
	if ok {
		// надо найти разницу в свойствах между базой и текущим
		//		sb := []string{base_word}
		for _, wp := range bw_prop {
			if wp.IsBase {
				rwp_b.Properties = append(rwp_b.Properties, wp.Properties...)
				//				sb = append(sb, wp.Properties...)
				break
			}
		}
		//		ssb := strings.Join(sb, "_")
		//		bb[ssb] = rwp_b
	}

	w_prop, ok := prop.WordFormPropertiesDict[word]
	if ok {
		for _, wp := range w_prop {
			rwp_l := ResultWordProp{}
			rwp_l.Word = word
			rwp_l.BaseWord = base_word
			//rwp_bc := ResultWordProp{}
			//copy(rwp_bc, rwp_b)
			//			sb := []string{word}
			if !wp.IsBase {
				rwp_l.Properties = append(rwp_b.Properties, wp.Properties...)
				//				sb = append(sb, rwp_l.Properties...)
				//				ssb := strings.Join(sb, "_")
				//				_, ok := bb[ssb]
				// fmt.Printf("! %v ok %v\r\n", ssb, ok)
				//				if !ok {
				rwp = append(rwp, rwp_l)
				//					bb[ssb] = rwp_l
				//				} else {
				//				}
			}
		}
	}
	return rwp
}

type HashWordProp struct {
	Value map[string]ResultWordProp
}

func InitHashWordProp() *HashWordProp {
	hwp := HashWordProp{}
	hwp.Value = make(map[string]ResultWordProp)
	return &hwp
}

func (hwp *HashWordProp) HashList(rwpl []ResultWordProp) []ResultWordProp {
	rwpln := []ResultWordProp{}
	for _, rwp := range rwpl {
		sb := []string{rwp.Word}
		sb = append(sb, rwp.Properties...)
		ssb := strings.Join(sb, "_")
		_, ok := hwp.Value[ssb]
		// fmt.Printf("! %v ok %v\r\n", ssb, ok)
		if ok {
		} else {
			rwpln = append(rwpln, rwp)
			hwp.Value[ssb] = rwp
		}
	}
	return rwpln
}

func FindWordCorpora1(word string, oc *OCorpora) (string, string) {
	wfdl := []string{word}
	ds, num, err := oc.Sdb.FindRecordIndexString([]string{"word_form"}, wfdl)
	if err != nil {
		fmt.Printf("Error %v\r\n", err)
		return "", ""
	}
	// fmt.Printf("ds %v num %v err %v\r\n", ds, num, err)

	if num == 0 {
		if len(ds) > 0 {
			res := ds[0].FieldsValue
			if len(res) > 0 {
				wsl := strings.Split(res[1], "_")
				return wsl[0], res[2]
			}
		}
	}
	return "", ""
}

func FindWordCorpora2(word string, oc *OCorpora) (string, string) {
	word_data1 := []string{word}
	ds1, num, err1 := oc.Sdb.FindRecordIndexString([]string{"word"}, word_data1)
	if err1 != nil {
		fmt.Printf("Error %v\r\n", err1)
		return "", ""
	}

	if num == 0 {
		// fmt.Printf("ds1 %v err1 %v\r\n", ds1, err1)
		if len(ds1) > 0 {
			for i := range ds1 {
				if len(ds1[i].FieldsValue) > 0 {
					res := ds1[i].FieldsValue
					if len(res) > 0 {
						//data := res[2]
						return res[1], res[2]
					}
				}
			}
		}
	}
	return "", ""
}

func (oc *OCorpora) TestWordCorporaShort1(co *Corpora) {
	for i := range co.WordProp {
		wp := co.WordProp[i]
		word_data := wp.BaseWordForm
		wp_data, err := json.Marshal(&wp)
		if err != nil {
			fmt.Println("error:", err)
			//return err
		}
		res1, res2 := FindWordCorpora2(word_data, oc)
		if res1 != "" {
			if res1 != word_data {
				fmt.Printf("error data word_data %v res1 %v\r\n", word_data, res1)
			}
			if res2 != string(wp_data) {
				fmt.Printf("error data word_data %v \r\n\tres1 %v \r\n\t string(wp_data) %v\r\n", word_data, res2, string(wp_data))
			}
		} else {
			fmt.Printf("error not find word_data %v\r\n", word_data)
		}
		for k, v := range wp.WordFormPropertiesDict {
			//word_form_data := k
			res1, res2 := FindWordCorpora1(k, oc)
			if res1 != "" {
				local_data, err := json.Marshal(&v)
				if err != nil {
					fmt.Println("error:", err)
					//return err
				}
				if res2 != string(local_data) {
					fmt.Printf("error data word_form_data %v \r\n\t res2 %v \r\n\t ld %v\r\n", k, res2, string(local_data))
				}
			} else {
				fmt.Printf("error not found word_form %v\r\n", k)
			}
		}
	}
}

func (oc *OCorpora) TestWordCorporaShort(co *Corpora) {
	for i := range co.WordProp {
		wp := co.WordProp[i]
		word_data := wp.BaseWordForm
		/*
		   wp_data, err := json.Marshal(&wp)
		   if err != nil {
		       fmt.Println("error:", err)
		       //return err
		   }
		*/
		/*
		           res1, res2 := FindWordCorpora2(word_data, oc)
		           if res1 != "" {
		               if res1 != word_data {
		                   fmt.Printf("error data word_data %v res1 %v\r\n", word_data, res1)
		               }
		               if res2 != string(wp_data) {
		   //                fmt.Printf("error data word_data %v \r\n\tres1 %v \r\n\t string(wp_data) %v\r\n", word_data, res2, string(wp_data))
		               }
		           } else {
		               fmt.Printf("error not find word_data %v\r\n", word_data)
		           }
		*/
		for k := range wp.WordFormPropertiesDict {
			//word_form_data := k
			result := FindWordCorporaShort(k, oc)
			if len(result) == 0 {
				fmt.Printf("error not found word_form %v\r\n", k)
			} else {
				n := len(result)
				for i := range result {
					if result[i][1] != word_data {
						n = n - 1
					}
				}
				if n == 0 {
					fmt.Printf("error not found base word_form %v\r\n", k)
				}
			}
			/*
			               res1, res2 := FindWordCorpora1(k, oc)
			               if res1 != "" {
			                   local_data, err := json.Marshal(&v)
			                   if err != nil {
			                       fmt.Println("error:", err)
			                       //return err
			                   }
			                   if res2 != string(local_data) {
			   //                    fmt.Printf("error data word_form_data %v \r\n\t res2 %v \r\n\t ld %v\r\n", k, res2, string(local_data))
			                   }
			               } else {
			                   fmt.Printf("error not found word_form %v\r\n", k)
			               }
			*/
		}
	}
}

func (oc *OCorpora) CloseSmallDB() {
	oc.Sdb.CloseData()
}
