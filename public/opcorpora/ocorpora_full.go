package opcorpora

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"

	"github.com/wanderer69/SmallDB/public/index"
)

func OpenOCorporaFull(path string, debug int) (*OCorpora, error) {
	sdb := index.InitSmallDB(path) // "./OCorporaDBFull"
	sdb.Debug = 0
	GlobalDictInit()
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
					attr, tag := Tag2StrAttrInt(wpi[j].Properties[m])
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
							attr, tag := Tag2StrAttrInt(wpi_n[l].Properties[n])
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

func (oc *OCorpora) FindWordCorporaFull(word string) ([]ResultWordProp, error) {
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
