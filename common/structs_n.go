package common

type WordItem struct {
	BaseWord string
	Properties map[string]string
}

type Command struct {
	Cmd             string			// команда
	Word			string			// слово для поиска
	ID     			string         	// идентификатор в очереди обработки
}

type CommandAnswer struct {
	Cmd         string			// команда
	Word		string
	Words		[][][]string	// найденные слова
	ID			string   		// идентификатор в очереди обработки
	Result      string
	Error       string
}

