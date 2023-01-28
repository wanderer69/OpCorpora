package op_corpora

var POS_dict map[string]string
var animacy_dict map[string]string
var genders_dict map[string]string
var numbers_dict map[string]string
var cases_dict map[string]string
var aspects_dict map[string]string
var transitivity_dict map[string]string
var persons_dict map[string]string
var tenses_dict map[string]string
var moods_dict map[string]string
var voices_dict map[string]string
var involvement_dict map[string]string
var attributes_dict map[string]string

var POS_dict_int map[string]string
var POS_array_int []string
var animacy_dict_int map[string]string
var genders_dict_int map[string]string
var numbers_dict_int map[string]string
var cases_dict_int map[string]string
var aspects_dict_int map[string]string
var transitivity_dict_int map[string]string
var persons_dict_int map[string]string
var tenses_dict_int map[string]string
var moods_dict_int map[string]string
var voices_dict_int map[string]string
var involvement_dict_int map[string]string

var POS_dict_back map[string]string
var animacy_dict_back map[string]string
var genders_dict_back map[string]string
var numbers_dict_back map[string]string
var cases_dict_back map[string]string
var aspects_dict_back map[string]string
var transitivity_dict_back map[string]string
var persons_dict_back map[string]string
var tenses_dict_back map[string]string
var moods_dict_back map[string]string
var voices_dict_back map[string]string
var involvement_dict_back map[string]string

var attributes_dict_int map[string]string
var attributes_dict_back map[string]string

func Global_Dict_Init() {
	POS_dict = make(map[string]string)
	POS_dict["NOUN"] = "имя существительное"
	POS_dict["ADJF"] = "имя прилагательное (полное)"
	POS_dict["ADJS"] = "имя прилагательное (краткое)"
	POS_dict["COMP"] = "компаратив"
	POS_dict["VERB"] = "глагол (личная форма)"
	POS_dict["INFN"] = "глагол (инфинитив)"
	POS_dict["PRTF"] = "причастие (полное)"
	POS_dict["PRTS"] = "причастие (краткое)"
	POS_dict["GRND"] = "деепричастие"
	POS_dict["NUMR"] = "числительное"
	POS_dict["ADVB"] = "наречие"
	POS_dict["NPRO"] = "местоимение-существительное"
	POS_dict["PRED"] = "предикатив"
	POS_dict["PREP"] = "предлог"
	POS_dict["CONJ"] = "союз"
	POS_dict["PRCL"] = "частица"
	POS_dict["INTJ"] = "междометие"
	POS_dict["PNCT"] = "знак пунктуации"
	POS_dict["LATN"] = "латинский"
	POS_dict["NUMB"] = "число"
	POS_dict["UNKN"] = "неопознанное"

	animacy_dict = make(map[string]string)
	animacy_dict["anim"] = "одушевлённое"
	animacy_dict["inan"] = "неодушевлённое"

	genders_dict = make(map[string]string)
	genders_dict["masc"] = "мужской род"
	genders_dict["femn"] = "женский род"
	genders_dict["neut"] = "средний род"

	numbers_dict = make(map[string]string)
	numbers_dict["sing"] = "единственное число"
	numbers_dict["plur"] = "множественное число"

	cases_dict = make(map[string]string)
	cases_dict["nomn"] = "именительный падеж"
	cases_dict["gent"] = "родительный падеж"
	cases_dict["datv"] = "дательный падеж"
	cases_dict["accs"] = "винительный падеж"
	cases_dict["ablt"] = "творительный падеж"
	cases_dict["loct"] = "предложный падеж"
	cases_dict["voct"] = "звательный падеж"
	cases_dict["gen1"] = "первый родительный падеж"
	cases_dict["gen2"] = "второй родительный (частичный) падеж"
	cases_dict["acc2"] = "второй винительный падеж"
	cases_dict["loc1"] = "первый предложный падеж"
	cases_dict["loc2"] = "второй предложный (местный) падеж"

	aspects_dict = make(map[string]string)
	aspects_dict["perf"] = "совершенный вид"
	aspects_dict["impf"] = "несовершенный вид"

	transitivity_dict = make(map[string]string)
	transitivity_dict["tran"] = "переходный"
	transitivity_dict["intr"] = "непереходный"

	persons_dict = make(map[string]string)
	persons_dict["1per"] = "1 лицо"
	persons_dict["2per"] = "2 лицо"
	persons_dict["3per"] = "3 лицо"

	tenses_dict = make(map[string]string)
	tenses_dict["pres"] = "настоящее время"
	tenses_dict["past"] = "прошедшее время"
	tenses_dict["futr"] = "будущее время"

	moods_dict = make(map[string]string)
	moods_dict["indc"] = "изъявительное наклонение"
	moods_dict["impr"] = "повелительное наклонение"

	voices_dict = make(map[string]string)
	voices_dict["actv"] = "действительный залог"
	voices_dict["pssv"] = "страдательный залог"

	involvement_dict = make(map[string]string)
	involvement_dict["incl"] = "говорящий включён в действие"
	involvement_dict["excl"] = "говорящий не включён в действие"

	attributes_dict = make(map[string]string)
	attributes_dict["Subx"] = "возможна субстантивация"
	attributes_dict["Supr"] = "превосходная степень"
	attributes_dict["Qual"] = "качественное"
	attributes_dict["Apro"] = "местоименное"
	attributes_dict["Anum"] = "порядковое"
	attributes_dict["Poss"] = "притяжательное"
	attributes_dict["V-ey"] = "форма на -ею"
	attributes_dict["V-oy"] = "форма на -ою"
	attributes_dict["Cmp2"] = "сравнительная степень на по-"
	attributes_dict["V-ej"] = "форма компаратива на -ей"

	attributes_dict["Infr"] = "разговорное"
	attributes_dict["Slng"] = "жаргонное"
	attributes_dict["Arch"] = "устаревшее"
	attributes_dict["Litr"] = "литературный вариант"
	attributes_dict["Erro"] = "опечатка"
	attributes_dict["Dist"] = "искажение"
	attributes_dict["Ques"] = "вопросительное"
	attributes_dict["Dmns"] = "указательное"
	attributes_dict["Prnt"] = "вводное слово"
	attributes_dict["V-be"] = "форма на -ье"
	attributes_dict["V-en"] = "форма на -енен"
	attributes_dict["V-ie"] = "форма на -и- (веселие, твердостию); отчество с -ие"
	attributes_dict["V-bi"] = "форма на -ьи"
	attributes_dict["Fimp"] = "деепричастие от глагола несовершенного вида"
	attributes_dict["Prdx"] = "может выступать в роли предикатива"
	attributes_dict["Coun"] = "счётная форма"
	attributes_dict["Coll"] = "собирательное числительное"
	attributes_dict["V-sh"] = "деепричастие на -ши"
	attributes_dict["Af-p"] = "форма после предлога"
	attributes_dict["Inmx"] = "может использоваться как одуш. / неодуш."
	attributes_dict["Vpre"] = "Вариант предлога ( со, подо, ...)"
	attributes_dict["Anph"] = "Анафорическое (местоимение)"
	attributes_dict["Init"] = "Инициал"
	attributes_dict["Adjx"] = "может выступать в роли прилагательного"
	attributes_dict["Ms-f"] = "колебание по роду (м/ж/с): кофе, вольво"
	attributes_dict["Hypo"] = "гипотетическая форма слова (победю, асфальтовее)"

	attributes_dict["Sgtm"] = "singularia tantum"
	attributes_dict["Pltm"] = "pluralia tantum"
	attributes_dict["Fixd"] = "неизменяемое"
	attributes_dict["CAse"] = "категория падежа"

	attributes_dict["GNdr"] = "род / род не выражен"

	//	attributes_dict[""] = ""
	//	attributes_dict[""] = ""

	POS_dict_int = make(map[string]string)
	POS_dict_int["NOUN"] = "имя_существительное"
	POS_dict_int["ADJF"] = "имя_прилагательное_полное"
	POS_dict_int["ADJS"] = "имя_прилагательное_краткое"
	POS_dict_int["COMP"] = "компаратив"
	POS_dict_int["VERB"] = "глагол_личная_форма"
	POS_dict_int["INFN"] = "глагол_инфинитив"
	POS_dict_int["PRTF"] = "причастие_полное"
	POS_dict_int["PRTS"] = "причастие_краткое"
	POS_dict_int["GRND"] = "деепричастие"
	POS_dict_int["NUMR"] = "числительное"
	POS_dict_int["ADVB"] = "наречие"
	POS_dict_int["NPRO"] = "местоимение-существительное"
	POS_dict_int["PRED"] = "предикатив"
	POS_dict_int["PREP"] = "предлог"
	POS_dict_int["ADP"] = "предлог"
	POS_dict_int["CONJ"] = "союз"
	POS_dict_int["PRCL"] = "частица"
	POS_dict_int["INTJ"] = "междометие"
	POS_dict_int["PNCT"] = "знак_пунктуации"
	POS_dict_int["PUNCT"] = "знак_пунктуации"
	POS_dict_int["LATN"] = "латинский"
	POS_dict_int["NUMB"] = "число"
	POS_dict_int["UNKN"] = "неопознанное"

        POS_array_int = []string{
             "имя_существительное",        
             "имя_прилагательное_полное",  
             "имя прилагательное_краткое", 
             "компаратив",                 
             "глагол_личная_форма",        
             "глагол_инфинитив",           
             "причастие_полное",           
             "причастие_краткое",          
             "деепричастие",               
             "числительное",               
             "наречие",                    
             "местоимение-существительное",
             "предикатив",                 
             "предлог",                    
             "союз",                       
             "частица",                    
             "междометие",                 
             "знак_пунктуации",            
             "латинский",                  
             "число",                      
             "неопознанное",               
        }

	animacy_dict_int = make(map[string]string)
	animacy_dict_int["anim"] = "одушевлённое"
	animacy_dict_int["inan"] = "неодушевлённое"

	genders_dict_int = make(map[string]string)
	genders_dict_int["masc"] = "мужской_род"
	genders_dict_int["femn"] = "женский_род"
	genders_dict_int["neut"] = "средний_род"

	numbers_dict_int = make(map[string]string)
	numbers_dict_int["sing"] = "единственное_число"
	numbers_dict_int["plur"] = "множественное_число"

	cases_dict_int = make(map[string]string)
	cases_dict_int["nomn"] = "именительный_падеж"
	cases_dict_int["gent"] = "родительный_падеж"
	cases_dict_int["datv"] = "дательный_падеж"
	cases_dict_int["accs"] = "винительный_падеж"
	cases_dict_int["ablt"] = "творительный_падеж"
	cases_dict_int["loct"] = "предложный_падеж"
	cases_dict_int["voct"] = "звательный_падеж"
	cases_dict_int["gen1"] = "первый_родительный_падеж"
	cases_dict_int["gen2"] = "второй_родительный_частичный_падеж"
	cases_dict_int["acc2"] = "второй_винительный_падеж"
	cases_dict_int["loc1"] = "первый_предложный_падеж"
	cases_dict_int["loc2"] = "второй_предложный_местный_падеж"

	aspects_dict_int = make(map[string]string)
	aspects_dict_int["perf"] = "совершенный_вид"
	aspects_dict_int["impf"] = "несовершенный_вид"

	transitivity_dict_int = make(map[string]string)
	transitivity_dict_int["tran"] = "переходный"
	transitivity_dict_int["intr"] = "непереходный"

	persons_dict_int = make(map[string]string)
	persons_dict_int["1per"] = "первое_лицо"
	persons_dict_int["2per"] = "второе_лицо"
	persons_dict_int["3per"] = "третье_лицо"

	tenses_dict_int = make(map[string]string)
	tenses_dict_int["pres"] = "настоящее_время"
	tenses_dict_int["past"] = "прошедшее_время"
	tenses_dict_int["futr"] = "будущее_время"

	moods_dict_int = make(map[string]string)
	moods_dict_int["indc"] = "изъявительное_наклонение"
	moods_dict_int["impr"] = "повелительное_наклонение"

	voices_dict_int = make(map[string]string)
	voices_dict_int["actv"] = "действительный_залог"
	voices_dict_int["pssv"] = "страдательный_залог"

	involvement_dict_int = make(map[string]string)
	involvement_dict_int["incl"] = "говорящий_включён_в_действие"
	involvement_dict_int["excl"] = "говорящий_не_включён_в_действие"

	attributes_dict_int = make(map[string]string)
	attributes_dict_int["Inmx"] = "может_использоваться_как_неодушевлённое"
	attributes_dict_int["Sgtm"] = "единственное_только"
	attributes_dict_int["Subx"] = "возможна субстантивация"
	attributes_dict_int["Supr"] = "превосходная степень"
	attributes_dict_int["Qual"] = "качественное"
	attributes_dict_int["Apro"] = "местоименное"
	attributes_dict_int["Anum"] = "порядковое"
	attributes_dict_int["Poss"] = "притяжательное"
	attributes_dict_int["V-ey"] = "форма на -ею"
	attributes_dict_int["V-oy"] = "форма на -ою"
	attributes_dict_int["Cmp2"] = "сравнительная степень на по-"
	attributes_dict_int["V-ej"] = "форма компаратива на -ей"

	attributes_dict_int["Infr"] = "разговорное"
	attributes_dict_int["Slng"] = "жаргонное"
	attributes_dict_int["Arch"] = "устаревшее"
	attributes_dict_int["Litr"] = "литературный вариант"
	attributes_dict_int["Erro"] = "опечатка"
	attributes_dict_int["Dist"] = "искажение"
	attributes_dict_int["Ques"] = "вопросительное"
	attributes_dict_int["Dmns"] = "указательное"
	attributes_dict_int["Prnt"] = "вводное слово"
	attributes_dict_int["V-be"] = "форма на -ье"
	attributes_dict_int["V-en"] = "форма на -енен"
	attributes_dict_int["V-ie"] = "форма на -и- (веселие, твердостию); отчество с -ие"
	attributes_dict_int["V-bi"] = "форма на -ьи"
	attributes_dict_int["Fimp"] = "деепричастие от глагола несовершенного вида"
	attributes_dict_int["Prdx"] = "может выступать в роли предикатива"
	attributes_dict_int["Coun"] = "счётная форма"
	attributes_dict_int["Coll"] = "собирательное числительное"
	attributes_dict_int["V-sh"] = "деепричастие на -ши"
	attributes_dict_int["Af-p"] = "форма после предлога"
	attributes_dict_int["Inmx"] = "может использоваться как одуш. / неодуш."
	attributes_dict_int["Vpre"] = "Вариант предлога ( со, подо, ...)"
	attributes_dict_int["Anph"] = "Анафорическое (местоимение)"
	attributes_dict_int["Init"] = "Инициал"
	attributes_dict_int["Adjx"] = "может выступать в роли прилагательного"
	attributes_dict_int["Ms-f"] = "колебание по роду (м/ж/с): кофе, вольво"
	attributes_dict_int["Hypo"] = "гипотетическая форма слова (победю, асфальтовее)"

//	attributes_dict_int["Sgtm"] = "singularia tantum"
	attributes_dict_int["Pltm"] = "pluralia tantum"
	attributes_dict_int["Fixd"] = "неизменяемое"
	attributes_dict_int["CAse"] = "категория падежа"

	attributes_dict_int["GNdr"] = "род / род не выражен"
	attributes_dict_int["ms-f"] = "общий род (м/ж)"

	attributes_dict_int["Impe"] = "безличный"
	attributes_dict_int["Impx"] = "возможно безличное употребление"
	attributes_dict_int["Mult"] = "многократный"
	attributes_dict_int["Refl"] = "возвратный"

	POS_dict_back = make(map[string]string)
	POS_dict_back["имя_существительное"] = "NOUN"
	POS_dict_back["имя_прилагательное_полное"] = "ADJF"
	POS_dict_back["имя прилагательное_краткое"] = "ADJS"
	POS_dict_back["компаратив"] = "COMP"
	POS_dict_back["глагол_личная_форма"] = "VERB"
	POS_dict_back["глагол_инфинитив"] = "INFN"
	POS_dict_back["причастие_полное"] = "PRTF"
	POS_dict_back["причастие_краткое"] = "PRTS"
	POS_dict_back["деепричастие"] = "GRND"
	POS_dict_back["числительное"] = "NUMR"
	POS_dict_back["наречие"] = "ADVB"
	POS_dict_back["местоимение-существительное"] = "NPRO"
	POS_dict_back["предикатив"] = "PRED"
	POS_dict_back["предлог"] = "PREP"
	POS_dict_back["союз"] = "CONJ"
	POS_dict_back["частица"] = "PRCL"
	POS_dict_back["междометие"] = "INTJ"
	POS_dict_back["знак_пунктуации"] = "PNCT"
	POS_dict_back["латинский"] = "LATN"
	POS_dict_back["число"] = "NUMB"
	POS_dict_back["неопознанное"] = "UNKN"

	animacy_dict_back = make(map[string]string)
	animacy_dict_back["одушевлённое"] = "anim"
	animacy_dict_back["неодушевлённое"] = "inan"

	genders_dict_back = make(map[string]string)
	genders_dict_back["мужской_род"] = "masc"
	genders_dict_back["женский_род"] = "femn"
	genders_dict_back["средний_род"] = "neut"

	numbers_dict_back = make(map[string]string)
	numbers_dict_back["единственное_число"] = "sing"
	numbers_dict_back["множественное_число"] = "plur"

	cases_dict_back = make(map[string]string)
	cases_dict_back["именительный_падеж"] = "nomn"
	cases_dict_back["родительный_падеж"] = "gent"
	cases_dict_back["дательный_падеж"] = "datv"
	cases_dict_back["винительный_падеж"] = "accs"
	cases_dict_back["творительный_падеж"] = "ablt"
	cases_dict_back["предложный_падеж"] = "loct"
	cases_dict_back["звательный_падеж"] = "voct"
	cases_dict_back["первый_родительный_падеж"] = "gen1"
	cases_dict_back["второй_родительный_частичный_падеж"] = "gen2"
	cases_dict_back["второй_винительный_падеж"] = "acc2"
	cases_dict_back["первый_предложный_падеж"] = "loc1"
	cases_dict_back["второй_предложный_местный_падеж"] = "loc2"

	aspects_dict_back = make(map[string]string)
	aspects_dict_back["совершенный_вид"] = "perf"
	aspects_dict_back["несовершенный_вид"] = "impf"

	transitivity_dict_back = make(map[string]string)
	transitivity_dict_back["переходный"] = "tran"
	transitivity_dict_back["непереходный"] = "intr"

	persons_dict_back = make(map[string]string)
	persons_dict_back["первое_лицо"] = "1per"
	persons_dict_back["второе_лицо"] = "2per"
	persons_dict_back["третье_лицо"] = "3per"

	tenses_dict_back = make(map[string]string)
	tenses_dict_back["настоящее_время"] = "pres"
	tenses_dict_back["прошедшее_время"] = "past"
	tenses_dict_back["будущее_время"] = "futr"

	moods_dict_back = make(map[string]string)
	moods_dict_back["изъявительное_наклонение"] = "indc"
	moods_dict_back["повелительное_наклонение"] = "impr"

	voices_dict_back = make(map[string]string)
	voices_dict_back["действительный_залог"] = "actv"
	voices_dict_back["страдательный_залог"] = "pssv"

	involvement_dict_back = make(map[string]string)
	involvement_dict_back["говорящий_включён_в_действие"] = "incl"
	involvement_dict_back["говорящий_не_включён_в_действие"] = "excl"

	attributes_dict_back = make(map[string]string)
	attributes_dict_back["может_использоваться_как_неодушевлённое"] = "Inmx"
	attributes_dict_back["единственное_только"] = "Sgtm"
	attributes_dict_back["возможна субстантивация"] = "Subx"
	attributes_dict_back["превосходная степень"] = "Supr"
	attributes_dict_back["качественное"] = "Qual"
	attributes_dict_back["местоименное"] = "Apro"
	attributes_dict_back["порядковое"] = "Anum"
	attributes_dict_back["притяжательное"] = "Poss"
	attributes_dict_back["форма на -ею"] = "V-ey"
	attributes_dict_back["форма на -ою"] = "V-oy"
	attributes_dict_back["сравнительная степень на по-"] = "Cmp2"
	attributes_dict_back["форма компаратива на -ей"] = "V-ej"

	attributes_dict_back["разговорное"] = "Infr"
	attributes_dict_back["жаргонное"] = "Slng"
	attributes_dict_back["устаревшее"] = "Arch"
	attributes_dict_back["литературный вариант"] = "Litr"
	attributes_dict_back["опечатка"] = "Erro"
	attributes_dict_back["искажение"] = "Dist"
	attributes_dict_back["вопросительное"] = "Ques"
	attributes_dict_back["указательное"] = "Dmns"
	attributes_dict_back["вводное слово"] = "Prnt"
	attributes_dict_back["форма на -ье"] = "V-be"
	attributes_dict_back["форма на -енен"] = "V-en"
	attributes_dict_back["форма на -и- (веселие, твердостию); отчество с -ие"] = "V-ie"
	attributes_dict_back["форма на -ьи"] = "V-bi"
	attributes_dict_back["деепричастие от глагола несовершенного вида"] = "Fimp"
	attributes_dict_back["может выступать в роли предикатива"] = "Prdx"
	attributes_dict_back["счётная форма"] = "Coun"
	attributes_dict_back["собирательное числительное"] = "Coll"
	attributes_dict_back["деепричастие на -ши"] = "V-sh"
	attributes_dict_back["форма после предлога"] = "Af-p"
	attributes_dict_back["может использоваться как одуш. / неодуш."] = "Inmx"
	attributes_dict_back["Вариант предлога ( со, подо, ...)"] = "Vpre"
	attributes_dict_back["Анафорическое (местоимение)"] = "Anph"
	attributes_dict_back["Инициал"] = "Init"
	attributes_dict_back["может выступать в роли прилагательного"] = "Adjx"
	attributes_dict_back["колебание по роду (м/ж/с): кофе, вольво"] = "Ms-f"
	attributes_dict_back["гипотетическая форма слова (победю, асфальтовее)"] = "Hypo"

	attributes_dict_back["singularia tantum"] =  "Sgtm"
	attributes_dict_back["pluralia tantum"  ] =  "Pltm"
	attributes_dict_back["неизменяемое"     ] =  "Fixd"
	attributes_dict_back["категория падежа" ] =  "CAse"
	attributes_dict_back["род / род не выражен"] = "GNdr"
	attributes_dict_back["общий род (м/ж)"] = "ms-f"

	attributes_dict_back["безличный"] = "Impe"
	attributes_dict_back["возможно безличное употребление"] = "Impx"
	attributes_dict_back["многократный"] = "Mult"
	attributes_dict_back["возвратный"] = "Refl"

}

func Tag2str_int(tag string) string {
	//var ret string
	if value_old, ok := POS_dict_int[tag]; ok {
		return value_old
	} else if value_old, ok := animacy_dict_int[tag]; ok {
		return value_old
	} else if value_old, ok := aspects_dict_int[tag]; ok {
		return value_old
	} else if value_old, ok := cases_dict_int[tag]; ok {
		return value_old
	} else if value_old, ok := genders_dict_int[tag]; ok {
		return value_old
	} else if value_old, ok := involvement_dict_int[tag]; ok {
		return value_old
	} else if value_old, ok := moods_dict_int[tag]; ok {
		return value_old
	} else if value_old, ok := numbers_dict_int[tag]; ok {
		return value_old
	} else if value_old, ok := persons_dict_int[tag]; ok {
		return value_old
	} else if value_old, ok := tenses_dict_int[tag]; ok {
		return value_old
	} else if value_old, ok := transitivity_dict_int[tag]; ok {
		return value_old
	} else if value_old, ok := voices_dict_int[tag]; ok {
		return value_old
	} else if value_old, ok := attributes_dict_int[tag]; ok {
		return value_old
	}
	return ""
}

func Tag2str_attr_int(tag string) (string, string) {
	//var ret string
	if value_old, ok := POS_dict_int[tag]; ok {
		return "POS", value_old
	} else if value_old, ok := animacy_dict_int[tag]; ok {
		return "animacy", value_old
	} else if value_old, ok := aspects_dict_int[tag]; ok {
		return "aspects", value_old
	} else if value_old, ok := cases_dict_int[tag]; ok {
		return "cases", value_old
	} else if value_old, ok := genders_dict_int[tag]; ok {
		return "genders", value_old
	} else if value_old, ok := involvement_dict_int[tag]; ok {
		return "involvement", value_old
	} else if value_old, ok := moods_dict_int[tag]; ok {
		return "moods", value_old
	} else if value_old, ok := numbers_dict_int[tag]; ok {
		return "numbers", value_old
	} else if value_old, ok := persons_dict_int[tag]; ok {
		return "persons", value_old
	} else if value_old, ok := tenses_dict_int[tag]; ok {
		return "tenses", value_old
	} else if value_old, ok := transitivity_dict_int[tag]; ok {
		return "transitivity", value_old
	} else if value_old, ok := voices_dict_int[tag]; ok {
		return "voices", value_old
	} else if value_old, ok := attributes_dict_int[tag]; ok {
		return "attributes", value_old
	}
	return "", ""
}

func Tag2num_array_int(tag string) (int, bool) {
	//var ret string
	for i, _ := range POS_array_int {
	      if POS_array_int[i] == tag {
                    return i, true
	      }
	}
	return -1, false
}
