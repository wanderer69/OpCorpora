package opcorpora

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/wanderer69/SmallDB/public/index"
)

func OpenYo(path string) (*OCorpora, error) {
	sdb := index.InitSmallDB(path)
	sdb.Debug = 0
	GlobalDictInit()
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
