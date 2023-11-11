package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wanderer69/OpCorpora/public/opcorpora"
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

	ss, err := opcorpora.LoadSettings("settings.json")
	if err != nil {
		fmt.Printf("LoadSettings %v\r\n", err)
		ss := opcorpora.InitSettings()
		opcorpora.SaveSettings("settings.json", ss)
		return
	}
	fmt.Printf("Settings %+v\r\n", ss)

	restrict := 10000000
	switch *ModePtr {
	case "create_gob":
		d := opcorpora.LoadDictionaryXML()
		opcorpora.StoreDictionary(d)

	case "create_corpora_gob_new":
		d := opcorpora.LoadDictionary()
		co, res := opcorpora.ConvertOpenCorpora2HashTableNew(&d, ss.RestrictionsList)
		if res {
			opcorpora.StoreCorpora(co)
		}

	case "convert_to_Yo_from_corpora_gob":
		co := opcorpora.LoadCorpora(*GobPtr)
		co.Restrict = restrict
		oc, _ := opcorpora.OpenYo(*BasePtr)
		oc.ConvertToYo(co)
		oc.CloseSmallDB()

	case "check_from_corpora_gob_full":
		co := opcorpora.LoadCorpora(*GobPtr)
		bwl, err1 := opcorpora.LoadBadWordsList()
		if err1 != nil {
			fmt.Printf("Error : %v\r\n", err1)
			return
		}

		co.Restrict = restrict

		oc, err := opcorpora.OpenOCorporaFull("./OCorporaCheck", 0)
		if err != nil {
			fmt.Printf("Error : %v\r\n", err)
			return
		}

		oc.StoreToSmallDBFull(co, bwl)
		oc.CloseSmallDB()

	case "create_from_corpora_gob_full":
		co := opcorpora.LoadCorpora(*GobPtr)
		bwl, err1 := opcorpora.LoadBadWordsList()
		if err1 != nil {
			fmt.Printf("Error : %v\r\n", err1)
			return
		}

		co.Restrict = restrict
		oc, err := opcorpora.OpenOCorporaFull(*BasePtr, 0)
		if err != nil {
			fmt.Printf("Error : %v\r\n", err)
			return
		}
		oc.StoreToSmallDBFull(co, bwl)
		oc.CloseSmallDB()

	case "create_from_corpora_short_gob":
		co := opcorpora.LoadCorpora(*GobPtr)
		oc, _ := opcorpora.OpenOCorporaShort("./OCorporaDBShort")
		oc.StoreToSmallDBShort(co)
		oc.CloseSmallDB()
	case "find_dict_gob":
		d := opcorpora.LoadDictionary()
		word_form_data := *ArgsPtr
		res := opcorpora.FindDictionary(&d, word_form_data)
		fmt.Printf("res %v\r\n", res)
	case "find_dict_gob_convert":
		d := opcorpora.LoadDictionary()
		word_form_data := *ArgsPtr
		res := opcorpora.FindDictionary(&d, word_form_data)
		fmt.Printf("res %v\r\n", res)
		co, res1 := opcorpora.ConvertOpenCorpora2HashTable(&d, ss.RestrictionsList, "./", word_form_data)
		if res1 {
			opcorpora.FindWord(word_form_data, co)
		}
	case "find_corpora_gob":
		co := opcorpora.LoadCorpora(*GobPtr)
		word_form_data := *ArgsPtr
		opcorpora.FindWord(word_form_data, co)
	case "find":
		oc, _ := opcorpora.OpenOCorpora(*BasePtr)
		word_form_data := *ArgsPtr
		opcorpora.FindWordCorpora(word_form_data, oc)
		oc.CloseSmallDB()
	case "test_find":
		co := opcorpora.LoadCorpora(*GobPtr)
		oc, _ := opcorpora.OpenOCorpora(*BasePtr)
		//oc.StoreToSmallDB(co)
		oc.TestFindSmallDB(co)
		oc.CloseSmallDB()
		//oc, _ := OpenOCorpora()
		//word_form_data := *ArgsPtr
		//FindWordCorpora(word_form_data, oc)
		//oc.CloseSmallDB()
	case "test_find_full":
		co := opcorpora.LoadCorpora(*GobPtr)
		co.Restrict = restrict
		oc, _ := opcorpora.OpenOCorporaFull(*BasePtr, 0)
		oc.TestFindSmallDBFull(co)
		oc.CloseSmallDB()
	case "find_short":
		oc, _ := opcorpora.OpenOCorporaShort("./OCorporaDBShort")
		word_form_data := *ArgsPtr
		result := opcorpora.FindWordCorporaShort(word_form_data, oc)
		if len(result) == 0 {
		} else {
			for i := range result {
				fmt.Printf("-> %v %v %v\r\n", result[i][0], result[i][1], result[i][2])
			}
		}
		oc.CloseSmallDB()
	case "find_short_short":
		oc, _ := opcorpora.OpenOCorporaShort("./OCorporaDBShort")
		word_form_data := *ArgsPtr
		result := opcorpora.FindWordCorporaShort(word_form_data, oc)
		if len(result) == 0 {
		} else {
			hwp := opcorpora.InitHashWordProp()
			for i := range result {
				word := result[i][0]
				// base_word := result[i][1]
				data := opcorpora.GetWordProperties(result[i][2])
				rwpl := opcorpora.GetWordFromWordProperties(word, data)
				rwpln := hwp.HashList(rwpl)
				if len(rwpln) > 0 {
					fmt.Printf("-> %+v \r\n", rwpln)
				}
			}

		}
		oc.CloseSmallDB()
	case "test_short":
		oc, _ := opcorpora.OpenOCorporaShort("./OCorporaDBShort")
		//word_form_data := *ArgsPtr
		co := opcorpora.LoadCorpora(*GobPtr)
		oc.TestWordCorporaShort(co)
		oc.CloseSmallDB()

	case "find_links":
		d := opcorpora.LoadDictionaryOC()
		link_type := "3"
		ll, err := opcorpora.FindDictionaryOC(d, link_type)
		if err != nil {
			fmt.Printf("%v\r\n", err)
		} else {
			for i := range ll {
				fmt.Printf("%v\r\n", ll[i])
			}
		}
	}
}
