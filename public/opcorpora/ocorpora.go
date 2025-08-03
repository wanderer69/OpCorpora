package opcorpora

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strings"

	"github.com/wanderer69/SmallDB/public/index"
)

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

func (oc *OCorpora) FindWordCorpora(word string) {
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

func (oc *OCorpora) FindWordCorporaService(word string) ([][][]string, error) {
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

func (oc *OCorpora) FindWordCorpora1(word string) (string, string) {
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

func (oc *OCorpora) FindWordCorpora2(word string) (string, string) {
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

func (oc *OCorpora) CloseSmallDB() {
	oc.Sdb.CloseData()
}
