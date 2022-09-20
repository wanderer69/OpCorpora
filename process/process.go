package process

import (
	"fmt"
	"time"
	"errors"

	. "github.com/wanderer69/OpCorpora/common"
	. "github.com/wanderer69/OpCorpora/settings"
	"math/rand"
	moc "github.com/wanderer69/OpCorpora" 	
)

func init() {
	Init_Unique_Value()
}

type SendReceiveItem struct {
	Type string // send, receive
	Word string
	Id   string
}

type SendReceiveAnswerItem struct {
	Type   string // send, receive
	Result string
	Error  string
	Id     string
	Word string
	Words		[][][]string	// найденные слова
}

type QueueItem struct {
	Type string // send, receive
	Id   string
	Word string
	Words		[][][]string	// найденные слова
	State string
	Result string
	Error  string	
	Next *QueueItem
}

type Queue struct {
	Head *QueueItem
	Last *QueueItem
}

func NewQueue() *Queue {
	q := Queue{}
	q.Head = nil
	q.Last = nil
	return &q
}

func (q *Queue) Add(qi *QueueItem) error {
	if q == nil {
		return errors.New("pointer to queue must be not nil")
	}
	if qi == nil {
		return errors.New("pointer to item must be not nil")
	}
	if q.Head == nil {
		q.Head = qi
	} else {
		q.Last.Next = qi
	}
	qi.Next = nil
	q.Last = qi
	return nil
}

func (q *Queue) DeleteFirst() error {
	if q == nil {
		return errors.New("pointer to queue must be not nil")
	}
	if q.Head == nil {
		return errors.New("queue empty")
	}
	if q.Head.Next == nil {
		q.Last = nil
	}
	q.Head = q.Head.Next
	return nil
}

func (q *Queue) Delete(qi *QueueItem) error {
	if q == nil {
		return errors.New("pointer to queue must be not nil")
	}
	if q.Head == nil {
		return errors.New("queue empty")
	}
	if qi == nil {
		return errors.New("pointer to item must be not nil")
	}
	qii := q.Head
	var qii_prev *QueueItem
	qii_prev = nil
	for {
		if qii == qi {
			if qii_prev == nil {
				q.Head = qii.Next
			} else {
				qii_prev.Next = qii.Next				
			}
		}
		
		if qii.Next == nil {
			break
		} else {
			qii = qii.Next
		}		
	}
	if q.Head != nil {
		if q.Head.Next == nil {
			q.Last = nil
		}
	}
	return nil
}

// на входе каналы - командный и для возврата результата
func Tasker(s *Settings, cmd_ch chan *Command, answer_ch chan *CommandAnswer, oc *moc.OCorpora) {
	// global_send_id := 1
	
	channel_in := make(chan SendReceiveItem)
	channel_out := make(chan SendReceiveAnswerItem)
	channel_exit := make(chan bool)
	// обмен с однопоточным клиентом
	send_recieve := func() {
		flag := false
		for {
			select {
			case sri := <-channel_in:
				//fmt.Printf("sri.Type %v\r\n", sri.Type)
				switch sri.Type {
				case "find":
					result_ := "OK"
					result, err := moc.FindWordCorporaService(oc, sri.Word)
					error_ := ""
					if err != nil {
						result_ = "Error"
						error_ = fmt.Sprintf("%v", err)
					}
					srai := SendReceiveAnswerItem{}
					srai.Type = sri.Type
					srai.Id = sri.Id
					srai.Result = result_
					srai.Error = error_
					srai.Word = sri.Word
					srai.Words = result
					channel_out <- srai
				}
			case <-channel_exit:
				flag = true
			}
			if flag {
				break
			}
		}
	}
	go send_recieve()
	q := NewQueue()
	q_ := NewQueue() // очередь выполненных заданий на оправку
	// в цикле ожидаем  приход команды и интерпретируем ее
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	state := 0
	next_state := 0
	n := 0
	var srai SendReceiveAnswerItem
	//msg_dict := make(map[string]*Message) 

	for {
		select {
		case <-ticker.C:
			n = n + 1
			if n == 60 {
				n = 0
				// ставим в очередь очередной запрос на прием
				qi := QueueItem{}
				qi.Type = "receive"
				qi.State = "wait"				
				q.Add(&qi)
				// отправляем периодический запрос на прием
				//v := t.String()
				//log.Println("write:", v)
			}
		case command_in := <-cmd_ch:
			switch command_in.Cmd {
			case "find":
				// формируем установку в очередь отправки
				qi := QueueItem{}
				qi.Type = command_in.Cmd
				qi.Word = command_in.Word
				qi.Id = Unique_Value(10)
				qi.State = "wait"				
				q.Add(&qi)
				// формируем ответ
				ca := CommandAnswer{}
				ca.Cmd = command_in.Cmd
				ca.Result = "OK"
				ca.ID = qi.Id
				answer_ch <- &ca
			case "check_find":
				// проверяем, что слово найдено
				qi := q.Head
				state := ""
				var words [][][]string
				flag_ := false
				if qi != nil {
					for {
						//fmt.Printf("command_in.MessageInID %v qi.Id %v\r\n", command_in.MessageInID, qi.Id)
						if command_in.ID == qi.Id {
							state = qi.State
							flag_ = true
							break
						}	
						if qi.Next == nil {
							qi = nil
							break
						} else {
							qi = qi.Next
						}
					}
				}
				if !flag_ {
					// проверяем очередь
					qi = q_.Head
					if qi != nil {
						for {
  						    //fmt.Printf("command_in.MessageInID %v qi.Id %v\r\n", command_in.MessageInID, qi.Id)
							if command_in.ID == qi.Id {
								state = qi.State
								words = qi.Words
								q_.Delete(qi)
								flag_ = true 
								break
							}	
							if qi.Next == nil {
								qi = nil
								break
							} else {
								qi = qi.Next
							}
						}
					}
				}
				// формируем ответ
				ca := CommandAnswer{}
				ca.Cmd = command_in.Cmd
				if !flag_ {
					ca.Result = "error"
					ca.Error = fmt.Sprintf("query with id %v not found", command_in.ID)
				} else {
					ca.Result = state
					ca.Words = words
				}
				answer_ch <- &ca
			case "set_filter":
				
			case "check_mail":
				// можем сразу выдать принятую почту
				
			case "get_mail":

			case "check":
				// метод проверки работы сервиса
				result := "Error"
				switch command_in.ID {
				case "Ping":
					result = "Pong"
				default:
					result = "OK"
				}
				ca := CommandAnswer{}
				ca.Cmd = command_in.Cmd
				ca.Result = result
				answer_ch <- &ca
			case "quit":
				channel_exit <- true
			}
		case srai = <-channel_out:
			// ответ 
			state = next_state
		}
		switch state {
		case 0:
			// нет загрузки почтового клиента. смотрим очередь
			if q.Head != nil {
			    //fmt.Printf("!!!\r\n")
				// есть, берем в работу
				qi := q.Head
				sri := SendReceiveItem{}
				sri.Type = qi.Type
				sri.Word = qi.Word
				channel_in <- sri
				qi.State = "in progress"				
				state = -1 
				next_state = 10
			}	
		case 10:
			// получили ответ обработчика
			qi := q.Head
			switch qi.Type {
				case "find":
					// обработка произошла и что делать?
					// передать в очередь обработанных
					qi := q.Head
					qi.State = "OK"								
					qi.Result = srai.Result
					qi.Error = srai.Error
					qi.Word = srai.Word					
					qi.Words = srai.Words
					q_.Add(qi)
					q.DeleteFirst()
			} 
			// можем переходить к следующему
			state = 0
		case 11:
		}
	}

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
