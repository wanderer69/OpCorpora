package opcorpora

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/wanderer69/SmallDB/public/index"
)

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

func (oc *OCorpora) FindWordCorporaShort(word string) [][]string {
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
