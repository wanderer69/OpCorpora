package opcorpora

var pOSDict map[string]string
var animacyDict map[string]string
var gendersDict map[string]string
var numbersDict map[string]string
var casesDict map[string]string
var aspectsDict map[string]string
var transitivityDict map[string]string
var personsDict map[string]string
var tensesDict map[string]string
var moodsDict map[string]string
var voicesDict map[string]string
var involvementDict map[string]string
var attributesDict map[string]string

var pOSDictInt map[string]string
var pOSArrayInt []string
var animacyDictInt map[string]string
var gendersDictInt map[string]string
var numbersDictInt map[string]string
var casesDictInt map[string]string
var aspectsDictInt map[string]string
var transitivityDictInt map[string]string
var personsDictInt map[string]string
var tensesDictInt map[string]string
var moodsDictInt map[string]string
var voicesDictInt map[string]string
var involvementDictInt map[string]string

var pOSDictBack map[string]string
var animacyDictBack map[string]string
var gendersDictBack map[string]string
var numbersDictBack map[string]string
var casesDictBack map[string]string
var aspectsDictBack map[string]string
var transitivityDictBack map[string]string
var personsDictBack map[string]string
var tensesDictBack map[string]string
var moodsDictBack map[string]string
var voicesDictBack map[string]string
var involvementDictBack map[string]string

var attributesDictInt map[string]string
var attributesDictBack map[string]string

func GlobalDictInit() {
	pOSDict = make(map[string]string)
	pOSDict["NOUN"] = "имя существительное"
	pOSDict["ADJF"] = "имя прилагательное (полное)"
	pOSDict["ADJS"] = "имя прилагательное (краткое)"
	pOSDict["COMP"] = "компаратив"
	pOSDict["VERB"] = "глагол (личная форма)"
	pOSDict["INFN"] = "глагол (инфинитив)"
	pOSDict["PRTF"] = "причастие (полное)"
	pOSDict["PRTS"] = "причастие (краткое)"
	pOSDict["GRND"] = "деепричастие"
	pOSDict["NUMR"] = "числительное"
	pOSDict["ADVB"] = "наречие"
	pOSDict["NPRO"] = "местоимение-существительное"
	pOSDict["PRED"] = "предикатив"
	pOSDict["PREP"] = "предлог"
	pOSDict["CONJ"] = "союз"
	pOSDict["PRCL"] = "частица"
	pOSDict["INTJ"] = "междометие"
	pOSDict["PNCT"] = "знак пунктуации"
	pOSDict["LATN"] = "латинский"
	pOSDict["NUMB"] = "число"
	pOSDict["UNKN"] = "неопознанное"

	animacyDict = make(map[string]string)
	animacyDict["anim"] = "одушевлённое"
	animacyDict["inan"] = "неодушевлённое"

	gendersDict = make(map[string]string)
	gendersDict["masc"] = "мужской род"
	gendersDict["femn"] = "женский род"
	gendersDict["fem"] = "женский род"
	gendersDict["neut"] = "средний род"

	numbersDict = make(map[string]string)
	numbersDict["sing"] = "единственное число"
	numbersDict["plur"] = "множественное число"

	casesDict = make(map[string]string)
	casesDict["nomn"] = "именительный падеж"
	casesDict["gent"] = "родительный падеж"
	casesDict["datv"] = "дательный падеж"
	casesDict["accs"] = "винительный падеж"
	casesDict["ablt"] = "творительный падеж"
	casesDict["loct"] = "предложный падеж"
	casesDict["voct"] = "звательный падеж"
	casesDict["gen1"] = "первый родительный падеж"
	casesDict["gen2"] = "второй родительный (частичный) падеж"
	casesDict["acc2"] = "второй винительный падеж"
	casesDict["loc1"] = "первый предложный падеж"
	casesDict["loc2"] = "второй предложный (местный) падеж"

	aspectsDict = make(map[string]string)
	aspectsDict["perf"] = "совершенный вид"
	aspectsDict["impf"] = "несовершенный вид"
	aspectsDict["imp"] = "несовершенный вид"

	transitivityDict = make(map[string]string)
	transitivityDict["tran"] = "переходный"
	transitivityDict["intr"] = "непереходный"

	personsDict = make(map[string]string)
	personsDict["1per"] = "1 лицо"
	personsDict["1"] = "1 лицо"
	personsDict["2per"] = "2 лицо"
	personsDict["2"] = "2 лицо"
	personsDict["3per"] = "3 лицо"
	personsDict["3"] = "3 лицо"

	tensesDict = make(map[string]string)
	tensesDict["pres"] = "настоящее время"
	tensesDict["past"] = "прошедшее время"
	tensesDict["futr"] = "будущее время"

	moodsDict = make(map[string]string)
	moodsDict["indc"] = "изъявительное наклонение"
	moodsDict["impr"] = "повелительное наклонение"

	voicesDict = make(map[string]string)
	voicesDict["actv"] = "действительный залог"
	voicesDict["pssv"] = "страдательный залог"

	involvementDict = make(map[string]string)
	involvementDict["incl"] = "говорящий включён в действие"
	involvementDict["excl"] = "говорящий не включён в действие"

	attributesDict = make(map[string]string)
	attributesDict["Subx"] = "возможна субстантивация"
	attributesDict["Supr"] = "превосходная степень"
	attributesDict["Qual"] = "качественное"
	attributesDict["Apro"] = "местоименное"
	attributesDict["Anum"] = "порядковое"
	attributesDict["Poss"] = "притяжательное"
	attributesDict["V-ey"] = "форма на -ею"
	attributesDict["V-oy"] = "форма на -ою"
	attributesDict["Cmp2"] = "сравнительная степень на по-"
	attributesDict["V-ej"] = "форма компаратива на -ей"

	attributesDict["Infr"] = "разговорное"
	attributesDict["Slng"] = "жаргонное"
	attributesDict["Arch"] = "устаревшее"
	attributesDict["Litr"] = "литературный вариант"
	attributesDict["Erro"] = "опечатка"
	attributesDict["Dist"] = "искажение"
	attributesDict["Ques"] = "вопросительное"
	attributesDict["Dmns"] = "указательное"
	attributesDict["Prnt"] = "вводное слово"
	attributesDict["V-be"] = "форма на -ье"
	attributesDict["V-en"] = "форма на -енен"
	attributesDict["V-ie"] = "форма на -и- (веселие, твердостию); отчество с -ие"
	attributesDict["V-bi"] = "форма на -ьи"
	attributesDict["Fimp"] = "деепричастие от глагола несовершенного вида"
	attributesDict["Prdx"] = "может выступать в роли предикатива"
	attributesDict["Coun"] = "счётная форма"
	attributesDict["Coll"] = "собирательное числительное"
	attributesDict["V-sh"] = "деепричастие на -ши"
	attributesDict["Af-p"] = "форма после предлога"
	attributesDict["Inmx"] = "может использоваться как одуш. / неодуш."
	attributesDict["Vpre"] = "Вариант предлога ( со, подо, ...)"
	attributesDict["Anph"] = "Анафорическое (местоимение)"
	attributesDict["Init"] = "Инициал"
	attributesDict["Adjx"] = "может выступать в роли прилагательного"
	attributesDict["Ms-f"] = "колебание по роду (м/ж/с): кофе, вольво"
	attributesDict["Hypo"] = "гипотетическая форма слова (победю, асфальтовее)"

	attributesDict["Sgtm"] = "singularia tantum"
	attributesDict["Pltm"] = "pluralia tantum"
	attributesDict["Fixd"] = "неизменяемое"
	attributesDict["CAse"] = "категория падежа"

	attributesDict["GNdr"] = "род / род не выражен"

	//	attributesDict[""] = ""
	//	attributesDict[""] = ""

	pOSDictInt = make(map[string]string)
	pOSDictInt["NOUN"] = "имя_существительное"
	pOSDictInt["ADJF"] = "имя_прилагательное_полное"
	pOSDictInt["ADJS"] = "имя_прилагательное_краткое"
	pOSDictInt["ADJ"] = "имя_прилагательное"
	pOSDictInt["DET"] = "имя_прилагательное_определитель" // детерменатив
	pOSDictInt["COMP"] = "компаратив"
	pOSDictInt["VERB"] = "глагол_личная_форма"
	pOSDictInt["INFN"] = "глагол_инфинитив"
	pOSDictInt["AUX"] = "глагол_вспомогательный"
	pOSDictInt["PRTF"] = "причастие_полное"
	pOSDictInt["PRTS"] = "причастие_краткое"
	pOSDictInt["GRND"] = "деепричастие"
	pOSDictInt["NUMR"] = "числительное"
	pOSDictInt["NUM"] = "числительное"
	pOSDictInt["ADVB"] = "наречие"
	pOSDictInt["ADV"] = "наречие"
	pOSDictInt["NPRO"] = "местоимение-существительное"
	pOSDictInt["PRON"] = "местоимение-существительное"
	pOSDictInt["PRED"] = "предикатив"
	pOSDictInt["PREP"] = "предлог"
	pOSDictInt["ADP"] = "предлог"
	pOSDictInt["CONJ"] = "союз"
	pOSDictInt["SCONJ"] = "союз_подчинительный"
	pOSDictInt["CCONJ"] = "союз_координирующий"
	pOSDictInt["PRCL"] = "частица"
	pOSDictInt["PART"] = "частица"
	pOSDictInt["INTJ"] = "междометие"
	pOSDictInt["PNCT"] = "знак_пунктуации"
	pOSDictInt["PUNCT"] = "знак_пунктуации"
	pOSDictInt["LATN"] = "латинский"
	pOSDictInt["NUMB"] = "число"
	pOSDictInt["UNKN"] = "неопознанное"

	pOSArrayInt = []string{
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
		"имя_прилагательное_определитель",
		"союз_подчинительный",
	}

	animacyDictInt = make(map[string]string)
	animacyDictInt["anim"] = "одушевлённое"
	animacyDictInt["inan"] = "неодушевлённое"

	gendersDictInt = make(map[string]string)
	gendersDictInt["masc"] = "мужской_род"
	gendersDictInt["femn"] = "женский_род"
	gendersDictInt["fem"] = "женский_род"
	gendersDictInt["neut"] = "средний_род"

	numbersDictInt = make(map[string]string)
	numbersDictInt["sing"] = "единственное_число"
	numbersDictInt["plur"] = "множественное_число"

	casesDictInt = make(map[string]string)
	casesDictInt["nomn"] = "именительный_падеж"
	casesDictInt["nom"] = "именительный_падеж"
	casesDictInt["gent"] = "родительный_падеж"
	casesDictInt["gen"] = "родительный_падеж"
	casesDictInt["datv"] = "дательный_падеж"
	casesDictInt["dat"] = "дательный_падеж"
	casesDictInt["accs"] = "винительный_падеж"
	casesDictInt["acc"] = "винительный_падеж"
	casesDictInt["ablt"] = "творительный_падеж"
	casesDictInt["ins"] = "творительный_падеж"
	casesDictInt["loct"] = "предложный_падеж"
	casesDictInt["loc"] = "предложный_падеж"
	casesDictInt["voct"] = "звательный_падеж"
	casesDictInt["gen1"] = "первый_родительный_падеж"
	casesDictInt["gen2"] = "второй_родительный_частичный_падеж"
	casesDictInt["acc2"] = "второй_винительный_падеж"
	casesDictInt["loc1"] = "первый_предложный_падеж"
	casesDictInt["loc2"] = "второй_предложный_местный_падеж"

	aspectsDictInt = make(map[string]string)
	aspectsDictInt["perf"] = "совершенный_вид"
	aspectsDictInt["impf"] = "несовершенный_вид"
	aspectsDictInt["imp"] = "несовершенный_вид"

	transitivityDictInt = make(map[string]string)
	transitivityDictInt["tran"] = "переходный"
	transitivityDictInt["intr"] = "непереходный"

	personsDictInt = make(map[string]string)
	personsDictInt["1per"] = "первое_лицо"
	personsDictInt["1"] = "первое_лицо"
	personsDictInt["2per"] = "второе_лицо"
	personsDictInt["2"] = "второе_лицо"
	personsDictInt["3per"] = "третье_лицо"
	personsDictInt["3"] = "третье_лицо"

	tensesDictInt = make(map[string]string)
	tensesDictInt["pres"] = "настоящее_время"
	tensesDictInt["past"] = "прошедшее_время"
	tensesDictInt["futr"] = "будущее_время"

	moodsDictInt = make(map[string]string)
	moodsDictInt["indc"] = "изъявительное_наклонение"
	moodsDictInt["ind"] = "изъявительное_наклонение"
	moodsDictInt["impr"] = "повелительное_наклонение"

	voicesDictInt = make(map[string]string)
	voicesDictInt["actv"] = "действительный_залог"
	voicesDictInt["act"] = "действительный_залог"
	voicesDictInt["pssv"] = "страдательный_залог"
	voicesDictInt["mid"] = "средний_залог"

	involvementDictInt = make(map[string]string)
	involvementDictInt["incl"] = "говорящий_включён_в_действие"
	involvementDictInt["excl"] = "говорящий_не_включён_в_действие"

	attributesDictInt = make(map[string]string)
	attributesDictInt["Inmx"] = "может_использоваться_как_неодушевлённое"
	attributesDictInt["Sgtm"] = "единственное_только"
	attributesDictInt["Subx"] = "возможна субстантивация"
	attributesDictInt["Supr"] = "превосходная степень"
	attributesDictInt["Qual"] = "качественное"
	attributesDictInt["Apro"] = "местоименное"
	attributesDictInt["Anum"] = "порядковое"
	attributesDictInt["Poss"] = "притяжательное"
	attributesDictInt["V-ey"] = "форма на -ею"
	attributesDictInt["V-oy"] = "форма на -ою"
	attributesDictInt["Cmp2"] = "сравнительная степень на по-"
	attributesDictInt["V-ej"] = "форма компаратива на -ей"

	attributesDictInt["Infr"] = "разговорное"
	attributesDictInt["Slng"] = "жаргонное"
	attributesDictInt["Arch"] = "устаревшее"
	attributesDictInt["Litr"] = "литературный вариант"
	attributesDictInt["Erro"] = "опечатка"
	attributesDictInt["Dist"] = "искажение"
	attributesDictInt["Ques"] = "вопросительное"
	attributesDictInt["Dmns"] = "указательное"
	attributesDictInt["Prnt"] = "вводное слово"
	attributesDictInt["V-be"] = "форма на -ье"
	attributesDictInt["V-en"] = "форма на -енен"
	attributesDictInt["V-ie"] = "форма на -и- (веселие, твердостию); отчество с -ие"
	attributesDictInt["V-bi"] = "форма на -ьи"
	attributesDictInt["Fimp"] = "деепричастие от глагола несовершенного вида"
	attributesDictInt["Prdx"] = "может выступать в роли предикатива"
	attributesDictInt["Coun"] = "счётная форма"
	attributesDictInt["Coll"] = "собирательное числительное"
	attributesDictInt["V-sh"] = "деепричастие на -ши"
	attributesDictInt["Af-p"] = "форма после предлога"
	attributesDictInt["Inmx"] = "может использоваться как одуш. / неодуш."
	attributesDictInt["Vpre"] = "Вариант предлога ( со, подо, ...)"
	attributesDictInt["Anph"] = "Анафорическое (местоимение)"
	attributesDictInt["Init"] = "Инициал"
	attributesDictInt["Adjx"] = "может выступать в роли прилагательного"
	attributesDictInt["Ms-f"] = "колебание по роду (м/ж/с): кофе, вольво"
	attributesDictInt["Hypo"] = "гипотетическая форма слова (победю, асфальтовее)"

	//	attributesDictInt["Sgtm"] = "singularia tantum"
	attributesDictInt["Pltm"] = "pluralia tantum"
	attributesDictInt["Fixd"] = "неизменяемое"
	attributesDictInt["CAse"] = "категория падежа"

	attributesDictInt["GNdr"] = "род / род не выражен"
	attributesDictInt["ms-f"] = "общий род (м/ж)"

	attributesDictInt["Impe"] = "безличный"
	attributesDictInt["Impx"] = "возможно безличное употребление"
	attributesDictInt["Mult"] = "многократный"
	attributesDictInt["Refl"] = "возвратный"

	attributesDictInt["fin"] = "спрягаемый"
	attributesDictInt["pos"] = "положительная степень"

	pOSDictBack = make(map[string]string)
	pOSDictBack["имя_существительное"] = "NOUN"
	pOSDictBack["имя_прилагательное_полное"] = "ADJF"
	pOSDictBack["имя прилагательное_краткое"] = "ADJS"
	pOSDictBack["компаратив"] = "COMP"
	pOSDictBack["глагол_личная_форма"] = "VERB"
	pOSDictBack["глагол_инфинитив"] = "INFN"
	pOSDictBack["причастие_полное"] = "PRTF"
	pOSDictBack["причастие_краткое"] = "PRTS"
	pOSDictBack["деепричастие"] = "GRND"
	pOSDictBack["числительное"] = "NUMR"
	pOSDictBack["наречие"] = "ADVB"
	pOSDictBack["местоимение-существительное"] = "NPRO"
	pOSDictBack["предикатив"] = "PRED"
	pOSDictBack["предлог"] = "PREP"
	pOSDictBack["союз"] = "CONJ"
	pOSDictBack["частица"] = "PRCL"
	pOSDictBack["междометие"] = "INTJ"
	pOSDictBack["знак_пунктуации"] = "PNCT"
	pOSDictBack["латинский"] = "LATN"
	pOSDictBack["число"] = "NUMB"
	pOSDictBack["неопознанное"] = "UNKN"

	animacyDictBack = make(map[string]string)
	animacyDictBack["одушевлённое"] = "anim"
	animacyDictBack["неодушевлённое"] = "inan"

	gendersDictBack = make(map[string]string)
	gendersDictBack["мужской_род"] = "masc"
	gendersDictBack["женский_род"] = "femn"
	gendersDictBack["средний_род"] = "neut"

	numbersDictBack = make(map[string]string)
	numbersDictBack["единственное_число"] = "sing"
	numbersDictBack["множественное_число"] = "plur"

	casesDictBack = make(map[string]string)
	casesDictBack["именительный_падеж"] = "nomn"
	casesDictBack["родительный_падеж"] = "gent"
	casesDictBack["дательный_падеж"] = "datv"
	casesDictBack["винительный_падеж"] = "accs"
	casesDictBack["творительный_падеж"] = "ablt"
	casesDictBack["предложный_падеж"] = "loct"
	casesDictBack["звательный_падеж"] = "voct"
	casesDictBack["первый_родительный_падеж"] = "gen1"
	casesDictBack["второй_родительный_частичный_падеж"] = "gen2"
	casesDictBack["второй_винительный_падеж"] = "acc2"
	casesDictBack["первый_предложный_падеж"] = "loc1"
	casesDictBack["второй_предложный_местный_падеж"] = "loc2"

	aspectsDictBack = make(map[string]string)
	aspectsDictBack["совершенный_вид"] = "perf"
	aspectsDictBack["несовершенный_вид"] = "impf"

	transitivityDictBack = make(map[string]string)
	transitivityDictBack["переходный"] = "tran"
	transitivityDictBack["непереходный"] = "intr"

	personsDictBack = make(map[string]string)
	personsDictBack["первое_лицо"] = "1per"
	personsDictBack["второе_лицо"] = "2per"
	personsDictBack["третье_лицо"] = "3per"

	tensesDictBack = make(map[string]string)
	tensesDictBack["настоящее_время"] = "pres"
	tensesDictBack["прошедшее_время"] = "past"
	tensesDictBack["будущее_время"] = "futr"

	moodsDictBack = make(map[string]string)
	moodsDictBack["изъявительное_наклонение"] = "indc"
	moodsDictBack["повелительное_наклонение"] = "impr"

	voicesDictBack = make(map[string]string)
	voicesDictBack["действительный_залог"] = "actv"
	voicesDictBack["страдательный_залог"] = "pssv"

	involvementDictBack = make(map[string]string)
	involvementDictBack["говорящий_включён_в_действие"] = "incl"
	involvementDictBack["говорящий_не_включён_в_действие"] = "excl"

	attributesDictBack = make(map[string]string)
	attributesDictBack["может_использоваться_как_неодушевлённое"] = "Inmx"
	attributesDictBack["единственное_только"] = "Sgtm"
	attributesDictBack["возможна субстантивация"] = "Subx"
	attributesDictBack["превосходная степень"] = "Supr"
	attributesDictBack["качественное"] = "Qual"
	attributesDictBack["местоименное"] = "Apro"
	attributesDictBack["порядковое"] = "Anum"
	attributesDictBack["притяжательное"] = "Poss"
	attributesDictBack["форма на -ею"] = "V-ey"
	attributesDictBack["форма на -ою"] = "V-oy"
	attributesDictBack["сравнительная степень на по-"] = "Cmp2"
	attributesDictBack["форма компаратива на -ей"] = "V-ej"

	attributesDictBack["разговорное"] = "Infr"
	attributesDictBack["жаргонное"] = "Slng"
	attributesDictBack["устаревшее"] = "Arch"
	attributesDictBack["литературный вариант"] = "Litr"
	attributesDictBack["опечатка"] = "Erro"
	attributesDictBack["искажение"] = "Dist"
	attributesDictBack["вопросительное"] = "Ques"
	attributesDictBack["указательное"] = "Dmns"
	attributesDictBack["вводное слово"] = "Prnt"
	attributesDictBack["форма на -ье"] = "V-be"
	attributesDictBack["форма на -енен"] = "V-en"
	attributesDictBack["форма на -и- (веселие, твердостию); отчество с -ие"] = "V-ie"
	attributesDictBack["форма на -ьи"] = "V-bi"
	attributesDictBack["деепричастие от глагола несовершенного вида"] = "Fimp"
	attributesDictBack["может выступать в роли предикатива"] = "Prdx"
	attributesDictBack["счётная форма"] = "Coun"
	attributesDictBack["собирательное числительное"] = "Coll"
	attributesDictBack["деепричастие на -ши"] = "V-sh"
	attributesDictBack["форма после предлога"] = "Af-p"
	attributesDictBack["может использоваться как одуш. / неодуш."] = "Inmx"
	attributesDictBack["Вариант предлога ( со, подо, ...)"] = "Vpre"
	attributesDictBack["Анафорическое (местоимение)"] = "Anph"
	attributesDictBack["Инициал"] = "Init"
	attributesDictBack["может выступать в роли прилагательного"] = "Adjx"
	attributesDictBack["колебание по роду (м/ж/с): кофе, вольво"] = "Ms-f"
	attributesDictBack["гипотетическая форма слова (победю, асфальтовее)"] = "Hypo"

	attributesDictBack["singularia tantum"] = "Sgtm"
	attributesDictBack["pluralia tantum"] = "Pltm"
	attributesDictBack["неизменяемое"] = "Fixd"
	attributesDictBack["категория падежа"] = "CAse"
	attributesDictBack["род / род не выражен"] = "GNdr"
	attributesDictBack["общий род (м/ж)"] = "ms-f"

	attributesDictBack["безличный"] = "Impe"
	attributesDictBack["возможно безличное употребление"] = "Impx"
	attributesDictBack["многократный"] = "Mult"
	attributesDictBack["возвратный"] = "Refl"

}

func Tag2str_int(tag string) string {
	//var ret string
	if valueOld, ok := pOSDictInt[tag]; ok {
		return valueOld
	} else if valueOld, ok := animacyDictInt[tag]; ok {
		return valueOld
	} else if valueOld, ok := aspectsDictInt[tag]; ok {
		return valueOld
	} else if valueOld, ok := casesDictInt[tag]; ok {
		return valueOld
	} else if valueOld, ok := gendersDictInt[tag]; ok {
		return valueOld
	} else if valueOld, ok := involvementDictInt[tag]; ok {
		return valueOld
	} else if valueOld, ok := moodsDictInt[tag]; ok {
		return valueOld
	} else if valueOld, ok := numbersDictInt[tag]; ok {
		return valueOld
	} else if valueOld, ok := personsDictInt[tag]; ok {
		return valueOld
	} else if valueOld, ok := tensesDictInt[tag]; ok {
		return valueOld
	} else if valueOld, ok := transitivityDictInt[tag]; ok {
		return valueOld
	} else if valueOld, ok := voicesDictInt[tag]; ok {
		return valueOld
	} else if valueOld, ok := attributesDictInt[tag]; ok {
		return valueOld
	}
	return ""
}

func Tag2StrAttrInt(tag string) (string, string) {
	//var ret string
	if valueOld, ok := pOSDictInt[tag]; ok {
		return "POS", valueOld
	} else if valueOld, ok := animacyDictInt[tag]; ok {
		return "animacy", valueOld
	} else if valueOld, ok := aspectsDictInt[tag]; ok {
		return "aspects", valueOld
	} else if valueOld, ok := casesDictInt[tag]; ok {
		return "cases", valueOld
	} else if valueOld, ok := gendersDictInt[tag]; ok {
		return "genders", valueOld
	} else if valueOld, ok := involvementDictInt[tag]; ok {
		return "involvement", valueOld
	} else if valueOld, ok := moodsDictInt[tag]; ok {
		return "moods", valueOld
	} else if valueOld, ok := numbersDictInt[tag]; ok {
		return "numbers", valueOld
	} else if valueOld, ok := personsDictInt[tag]; ok {
		return "persons", valueOld
	} else if valueOld, ok := tensesDictInt[tag]; ok {
		return "tenses", valueOld
	} else if valueOld, ok := transitivityDictInt[tag]; ok {
		return "transitivity", valueOld
	} else if valueOld, ok := voicesDictInt[tag]; ok {
		return "voices", valueOld
	} else if valueOld, ok := attributesDictInt[tag]; ok {
		return "attributes", valueOld
	}
	return "", ""
}

func Tag2StrAttrRuInt(tag string) (string, string) {
	//var ret string
	if valueOld, ok := pOSDictInt[tag]; ok {
		return "часть_речи", valueOld
	} else if valueOld, ok := animacyDictInt[tag]; ok {
		return "одушевлённость", valueOld
	} else if valueOld, ok := aspectsDictInt[tag]; ok {
		return "совершённость", valueOld
	} else if valueOld, ok := casesDictInt[tag]; ok {
		return "падеж", valueOld
	} else if valueOld, ok := gendersDictInt[tag]; ok {
		return "пол", valueOld
	} else if valueOld, ok := involvementDictInt[tag]; ok {
		return "involvement", valueOld
	} else if valueOld, ok := moodsDictInt[tag]; ok {
		return "наклонение", valueOld
	} else if valueOld, ok := numbersDictInt[tag]; ok {
		return "число", valueOld
	} else if valueOld, ok := personsDictInt[tag]; ok {
		return "лицо", valueOld
	} else if valueOld, ok := tensesDictInt[tag]; ok {
		return "время", valueOld
	} else if valueOld, ok := transitivityDictInt[tag]; ok {
		return "переходность", valueOld
	} else if valueOld, ok := voicesDictInt[tag]; ok {
		return "залог", valueOld
	} else if valueOld, ok := attributesDictInt[tag]; ok {
		return "атрибут", valueOld
	}
	return "", ""
}

func Tag2NumArrayInt(tag string) (int, bool) {
	//var ret string
	for i := range pOSArrayInt {
		if pOSArrayInt[i] == tag {
			return i, true
		}
	}
	return -1, false
}
