package main

import (
	"flag"
	"fmt"
	"os"

	. "github.com/wanderer69/OpCorpora"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		fmt.Printf("_ [<base>] <mode> [<gob>] [<args>]\r\n")
		return
	}
	BasePtr := flag.String("base", "./WithoutYo2WithYo", "name base")
	ModePtr := flag.String("mode", "", "mode: convert")
	GobPtr := flag.String("gob", "../utils/opcorpora.gob", "gob")
	ArgsPtr := flag.String("args", "", "arg list")
	flag.Parse()

	fmt.Printf("base %v\r\n", *BasePtr)

	ss, err := LoadSettings("settings.json")
	if err != nil {
		fmt.Printf("LoadSettings %v\r\n", err)
		ss := InitSettings()
		SaveSettings("settings.json", ss)
		return
	}
	fmt.Printf("Settings %+v\r\n", ss)

	restrict := 10000000
	switch *ModePtr {
	case "create_gob":
		d := LoadDictionaryXML()
		StoreDictionary(d)

	case "create_corpora_gob_new":
		d := LoadDictionary()
		co, res := ConvertOpenCorpora2HashTableNew(&d, ss.RestrictionsList)
		if res {
			StoreCorpora(co)
		}

	case "convert_to_Yo_from_corpora_gob":
		co := LoadCorpora(*GobPtr)
		co.Restrict = restrict
		oc, _ := OpenYo(*BasePtr)
		oc.ConvertToYo(co)
		oc.CloseSmallDB()

	case "check_from_corpora_gob_full":
		co := LoadCorpora(*GobPtr)
		bwl, err1 := LoadBadWordsList()
		if err1 != nil {
			fmt.Printf("Error : %v\r\n", err1)
			return
		}

		co.Restrict = restrict

		oc, err := OpenOCorporaFull("./OCorporaCheck")
		if err != nil {
			fmt.Printf("Error : %v\r\n", err)
			return
		}

		oc.StoreToSmallDBFull(co, bwl)
		oc.CloseSmallDB()

	case "create_from_corpora_gob_full":
		co := LoadCorpora(*GobPtr)
		bwl, err1 := LoadBadWordsList()
		if err1 != nil {
			fmt.Printf("Error : %v\r\n", err1)
			return
		}

		co.Restrict = restrict
		oc, err := OpenOCorporaFull(*BasePtr)
		if err != nil {
			fmt.Printf("Error : %v\r\n", err)
			return
		}
		oc.StoreToSmallDBFull(co, bwl)
		oc.CloseSmallDB()

	case "create_from_corpora_short_gob":
		co := LoadCorpora(*GobPtr)
		oc, _ := OpenOCorporaShort("./OCorporaDBShort")
		oc.StoreToSmallDBShort(co)
		oc.CloseSmallDB()
	case "find_dict_gob":
		d := LoadDictionary()
		word_form_data := *ArgsPtr
		res := FindDictionary(&d, word_form_data)
		fmt.Printf("res %v\r\n", res)
	case "find_dict_gob_convert":
		d := LoadDictionary()
		word_form_data := *ArgsPtr
		res := FindDictionary(&d, word_form_data)
		fmt.Printf("res %v\r\n", res)
		co, res1 := ConvertOpenCorpora2HashTable(&d, ss.RestrictionsList, "./", word_form_data)
		if res1 {
			FindWord(word_form_data, co)
		}
	case "find_corpora_gob":
		co := LoadCorpora(*GobPtr)
		word_form_data := *ArgsPtr
		FindWord(word_form_data, co)
	case "find":
		oc, _ := OpenOCorpora(*BasePtr)
		word_form_data := *ArgsPtr
		FindWordCorpora(word_form_data, oc)
		oc.CloseSmallDB()
	case "test_find":
		co := LoadCorpora(*GobPtr)
		oc, _ := OpenOCorpora(*BasePtr)
		//oc.StoreToSmallDB(co)
		oc.TestFindSmallDB(co)
		oc.CloseSmallDB()
		//oc, _ := OpenOCorpora()
		//word_form_data := *ArgsPtr
		//FindWordCorpora(word_form_data, oc)
		//oc.CloseSmallDB()
	case "test_find_full":
		co := LoadCorpora(*GobPtr)
		co.Restrict = restrict
		oc, _ := OpenOCorporaFull(*BasePtr)
		oc.TestFindSmallDBFull(co)
		oc.CloseSmallDB()
	case "find_short":
		oc, _ := OpenOCorporaShort("./OCorporaDBShort")
		word_form_data := *ArgsPtr
		result := FindWordCorporaShort(word_form_data, oc)
		if len(result) == 0 {
		} else {
			for i, _ := range result {
				fmt.Printf("-> %v %v %v\r\n", result[i][0], result[i][1], result[i][2])
			}
		}
		oc.CloseSmallDB()
	case "find_short_short":
		oc, _ := OpenOCorporaShort("./OCorporaDBShort")
		word_form_data := *ArgsPtr
		result := FindWordCorporaShort(word_form_data, oc)
		if len(result) == 0 {
		} else {
			hwp := InitHashWordProp()
			for i, _ := range result {
				word := result[i][0]
				// base_word := result[i][1]
				data := GetWordProperties(result[i][2])
				rwpl := GetWordFromWordProperties(word, data)
				rwpln := hwp.HashList(rwpl)
				if len(rwpln) > 0 {
					fmt.Printf("-> %+v \r\n", rwpln)
				}
			}

		}
		oc.CloseSmallDB()
	case "test_short":
		oc, _ := OpenOCorporaShort("./OCorporaDBShort")
		//word_form_data := *ArgsPtr
		co := LoadCorpora(*GobPtr)
		oc.TestWordCorporaShort(co)
		oc.CloseSmallDB()

	case "find_links":
		d := LoadDictionaryOC()
		link_type := "3"
		ll, err := FindDictionaryOC(d, link_type)
		if err != nil {
			fmt.Printf("%v\r\n", err)
		} else {
			for i, _ := range ll {
				fmt.Printf("%v\r\n", ll[i])
			}
		}
	}
}
