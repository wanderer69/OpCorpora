package opcorpora

import "strings"

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
