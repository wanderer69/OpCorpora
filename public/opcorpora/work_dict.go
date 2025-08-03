package opcorpora

import (
	"bytes"
	"encoding/gob"
	"encoding/json"

	"fmt"

	"math/rand"
	"os"
	"strings"
	"time"
)

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
		word_data := wp.BaseWordForm
		flag_m := 0
		for k := range wp.WordFormPropertiesDict {
			word_form_data := k
			wpi := wpd[word_form_data]

			wfdl := []string{word_form_data}
			ds, num, err := oc.Sdb.FindRecordIndexString([]string{"word_form"}, wfdl)
			if err != nil {
				fmt.Printf("Error 1 %v num %v %v\r\n", err, num, wfdl)
				continue
			}
			if num != 0 {
				fmt.Printf("Error 2 %v %v\r\n", num, wfdl)
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
				flag := 0
				for n_rec := range ds {
					rec := ds[n_rec]
					val, err := oc.Sdb.GetFieldValueByName(rec, "word")
					if err != nil {
						fmt.Printf("Error 4! %v\r\n", err)
						return
					} else {
						if val == word_data {
							flag_p := 0
							cmp := 0
							for m := range wpi[j].Properties {
								attr, tag := Tag2StrAttrInt(wpi[j].Properties[m])
								if len(attr) != 0 {
									/*
										if ((((((((((attr == "POS" && attr == "animacy") && attr == "aspects") && attr == "cases") && attr == "genders") &&
											attr == "involvement") && attr == "moods") && attr == "numbers") && attr == "persons") && attr == "tenses") &&
											attr == "transitivity") && attr != "voices" {
									*/
									if attr == "POS" && attr == "animacy" && attr == "aspects" && attr == "cases" && attr == "genders" &&
										attr == "involvement" && attr == "moods" && attr == "numbers" && attr == "persons" && attr == "tenses" &&
										attr == "transitivity" && attr != "voices" {
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
		word_data := wp.BaseWordForm

		for k := range wp.WordFormPropertiesDict {
			word_form_data := k
			fmt.Printf("word form %v word %v\r\n", word_form_data, word_data)
			oc.FindWordCorporaTest(word_form_data, word_data)
		}
		if i == 1000 {
			break
		}
	}
}

func (oc *OCorpora) FindWordCorporaTest(word string, word_data string) {
	wfdl := []string{word}
	ds, num, err := oc.Sdb.FindRecordIndexString([]string{"word_form"}, wfdl)
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

func GetWordProperties(data string) WordProperties {
	var prop WordProperties
	// unmarshall it
	err1 := json.Unmarshal([]byte(data), &prop)
	if err1 != nil {
		fmt.Println("error:", err1)
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
	bw_prop, ok := prop.WordFormPropertiesDict[base_word]
	if ok {
		// надо найти разницу в свойствах между базой и текущим
		for _, wp := range bw_prop {
			if wp.IsBase {
				rwp_b.Properties = append(rwp_b.Properties, wp.Properties...)
				break
			}
		}
	}

	w_prop, ok := prop.WordFormPropertiesDict[word]
	if ok {
		for _, wp := range w_prop {
			rwp_l := ResultWordProp{}
			rwp_l.Word = word
			rwp_l.BaseWord = base_word
			if !wp.IsBase {
				rwp_l.Properties = append(rwp_b.Properties, wp.Properties...)
				rwp = append(rwp, rwp_l)
			}
		}
	}
	return rwp
}

func (oc *OCorpora) TestWordCorporaShort1(co *Corpora) {
	for i := range co.WordProp {
		wp := co.WordProp[i]
		word_data := wp.BaseWordForm
		wp_data, err := json.Marshal(&wp)
		if err != nil {
			fmt.Println("error:", err)
		}
		res1, res2 := oc.FindWordCorpora2(word_data)
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
			res1, res2 := oc.FindWordCorpora1(k)
			if res1 != "" {
				local_data, err := json.Marshal(&v)
				if err != nil {
					fmt.Println("error:", err)
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
		for k := range wp.WordFormPropertiesDict {
			result := oc.FindWordCorporaShort(k)
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
		}
	}
}
