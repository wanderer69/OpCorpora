package gateway

import (
	"errors"
	"fmt"
	"net"

	proc "github.com/wanderer69/OpCorpora/internal/process"
	"github.com/wanderer69/OpCorpora/internal/settings"
	"github.com/wanderer69/OpCorpora/pkg/proto"
	"github.com/wanderer69/OpCorpora/public/common"
	moc "github.com/wanderer69/OpCorpora/public/opcorpora"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type server struct {
	proto.UnimplementedOpCorporaServiceServer

	oc       *moc.OCorpora
	s        *settings.Settings
	cmdCh    chan *common.Command
	answerCh chan *common.CommandAnswer

	execCmd func() (*common.CommandAnswer, error)

	/*
		mode      func(context.Context, *proto.ModeRequest) (*proto.ModeResponse, error)
		check     func(context.Context, *proto.CheckRequest) (*proto.CheckResponse, error)
		findWord  func(context.Context, *proto.FindWordRequest) (*proto.FindWordResponse, error)
		checkFind func(context.Context, *proto.CheckFindRequest) (*proto.CheckFindResponse, error)
		stat      func(context.Context, *proto.StatRequest) (*proto.StatResponse, error)
	*/
	/*
			Mode(context.Context, *ModeRequest) (*ModeResponse, error)
		    FindWord(context.Context, *FindWordRequest) (*FindWordResponse, error)
		    Check(context.Context, *CheckRequest) (*CheckResponse, error)
		    CheckFind(context.Context, *CheckFindRequest) (*CheckFindResponse, error)
		    Stat(context.Context, *StatRequest) (*StatResponse, error)
	*/
}

func (s *server) Mode(ctx context.Context, request *proto.ModeRequest) (*proto.ModeResponse, error) {
	// fmt.Printf("%v\r\n", request.Query)
	result := "Error"
	switch request.Mode {
	case "sync":
		//flag_mode = request.Mode
		result = "OK"
	case "async":
		//flag_mode = request.Mode
		result = "OK"
	}
	response := &proto.ModeResponse{
		Result: result,
	}
	return response, nil
	/*
	   r, err := s.mode(ctx, request)

	   return r, err
	*/
}

func (s *server) Check(ctx context.Context, request *proto.CheckRequest) (*proto.CheckResponse, error) {
	// fmt.Printf("%v\r\n", request.Query)
	switch s.s.Mode {
	case "sync":
	case "async":
		cmd := common.Command{}
		cmd.Cmd = "check"
		cmd.ID = request.Query
		s.cmdCh <- &cmd
		//var ca *common.CommandAnswer
		ca, err := s.execCmd()
		result := ""
		if err != nil {
			result = fmt.Sprintf("Error %v", err)
		} else {
			if ca == nil {
				result = "error ca == nil!"
			} else {
				result = ca.Result + " " + ca.Error
			}
		}
		response := &proto.CheckResponse{
			Result: result,
		}
		return response, nil
	}
	return nil, errors.New("bad mode")
	/*
	   r, err := s.check(ctx, request)
	   return r, err
	*/
}

func (s *server) FindWord(ctx context.Context, request *proto.FindWordRequest) (*proto.FindWordResponse, error) {
	//fmt.Printf("%v\r\n", request.Token)
	switch s.s.Mode {
	case "sync":
		result, err := moc.FindWordCorporaService(s.oc, request.Word)
		if err != nil {
			response := &proto.FindWordResponse{
				Result: "Error",
			}
			return response, nil
		}
		wra := []*proto.WordRecord{}
		for i := range result {
			wr := proto.WordRecord{}
			for j := range result[i] {
				if len(result[i][j][1]) > 0 {
					wp := proto.WordProperty{Property: result[i][j][0], Value: result[i][j][1]}
					wr.WordProperties = append(wr.WordProperties, &wp)
				}
			}
			wr.Word = request.Word
			wra = append(wra, &wr)
		}
		response := &proto.FindWordResponse{
			Result:      "OK",
			WordRecords: wra,
		}

		return response, nil
	case "async":
		cmd := common.Command{}
		cmd.Cmd = "find"
		cmd.Word = request.Word

		s.cmdCh <- &cmd

		ca, err := s.execCmd()
		result := ""
		errorValue := ""
		if err != nil {
			result = "Error"
			errorValue = fmt.Sprintf("%v", err)
		} else {
			result = ca.Result
			errorValue = ca.Error
		}

		response := &proto.FindWordResponse{
			Result: result,
			Error:  &errorValue,
			ReqId:  &ca.ID,
		}

		return response, nil

	}
	return nil, errors.New("bad mode")

	/*
	   r, err := s.findWord(ctx, request)
	   return r, err
	*/
}

func (s *server) CheckFind(ctx context.Context, request *proto.CheckFindRequest) (*proto.CheckFindResponse, error) {
	//fmt.Printf("%v\r\n", request.ReqId)
	switch s.s.Mode {
	case "sync":

	case "async":
		cmd := common.Command{}
		cmd.Cmd = "check_find"
		cmd.ID = request.ReqId

		s.cmdCh <- &cmd

		ca, err := s.execCmd()
		result := "OK"
		var errorValue *string
		var wra []*proto.WordRecord

		if err != nil {
			result = "Error"
			e := fmt.Sprintf("%v", err)
			errorValue = &e
		} else {
			result = ca.Result
			errorValue = &ca.Error
			if result == "OK" {
				wra = []*proto.WordRecord{}
				for i := range ca.Words {
					wr := proto.WordRecord{}
					for j := range ca.Words[i] {
						if len(ca.Words[i][j][1]) > 0 {
							wp := proto.WordProperty{Property: ca.Words[i][j][0], Value: ca.Words[i][j][1]}
							wr.WordProperties = append(wr.WordProperties, &wp)
						}
					}
					wr.Word = ca.Word
					wra = append(wra, &wr)
				}
			}
		}

		response := &proto.CheckFindResponse{
			Result:      result,
			Error:       errorValue,
			WordRecords: wra,
		}
		return response, nil
	}
	return nil, errors.New("bad mode")
	/*
	   r, err := s.checkFind(ctx, request)
	   return r, err
	*/
}

func (s *server) Stat(ctx context.Context, in *proto.StatRequest) (*proto.StatResponse, error) {
	result := ""
	response := &proto.StatResponse{
		Result: result,
	}
	return response, nil
	/*
	   r, err := s.stat(ctx, in)
	   return r, err
	*/
}

/*
func (s *server) mustEmbedUnimplementedOpCorporaServiceServer() {

}
*/

/*
type ModeFunc func(ctx context.Context, mode string) (string, error)
type CheckFunc func(ctx context.Context, query string) (string, error)
type FindWordFunc func(ctx context.Context, word string) (string, string, string, map[string]string, error)
type CheckFindFunc func(ctx context.Context, id string) (string, string, string, map[string]string, error)
type StatFunc func(ctx context.Context, mode string) (string, string, string, map[string]string, error)
*/

func GRPCServer(s *settings.Settings) error {
	ss := fmt.Sprintf(":%v", s.PortClient)
	listener, err := net.Listen("tcp", ss)

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
		return err
	}

	var cmdCh chan *common.Command
	var answerCh chan *common.CommandAnswer

	oc, err := moc.OpenOCorporaFull(s.OpCorporaPath, 0)
	if err != nil {
		grpclog.Fatalf("failed to open db: %v", err)
		return err
	}

	initAsync := func(oc *moc.OCorpora) {
		cmdCh = make(chan *common.Command)
		answerCh = make(chan *common.CommandAnswer)

		go proc.Tasker(s, cmdCh, answerCh, oc)
	}
	execCmd := func() (*common.CommandAnswer, error) {
		// надо дождаться ответа
		var ca *common.CommandAnswer
		for ca = range answerCh {
			break
		}
		return ca, nil
	}

	switch s.Mode {
	case "sync":

	case "async":
		initAsync(oc)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	serverItem := server{
		/*
			check: func(c context.Context, request *proto.CheckRequest) (*proto.CheckResponse, error) {
				// fmt.Printf("%v\r\n", request.Query)
				switch s.Mode {
				case "sync":
				case "async":
					cmd := common.Command{}
					cmd.Cmd = "check"
					cmd.ID = request.Query
					cmdCh <- &cmd
					//var ca *common.CommandAnswer
					ca, err := execCmd()
					result := ""
					if err != nil {
						result = fmt.Sprintf("Error %v", err)
					} else {
						if ca == nil {
							result = "error ca == nil!"
						} else {
							result = ca.Result + " " + ca.Error
						}
					}
					response := &proto.CheckResponse{
						Result: result,
					}
					return response, nil
				}
				return nil, errors.New("bad mode")
			},
			mode: func(c context.Context, request *proto.ModeRequest) (*proto.ModeResponse, error) {
				// fmt.Printf("%v\r\n", request.Query)
				result := "Error"
				switch request.Mode {
				case "sync":
					//flag_mode = request.Mode
					result = "OK"
				case "async":
					//flag_mode = request.Mode
					result = "OK"
				}
				response := &proto.ModeResponse{
					Result: result,
				}
				return response, nil
			},
			findWord: func(c context.Context, request *proto.FindWordRequest) (*proto.FindWordResponse, error) {
				//fmt.Printf("%v\r\n", request.Token)
				switch s.Mode {
				case "sync":
					result, err := moc.FindWordCorporaService(oc, request.Word)
					if err != nil {
						response := &proto.FindWordResponse{
							Result: "Error",
						}
						return response, nil
					}
					wra := []*proto.WordRecord{}
					for i := range result {
						wr := proto.WordRecord{}
						for j := range result[i] {
							if len(result[i][j][1]) > 0 {
								wp := proto.WordProperty{Property: result[i][j][0], Value: result[i][j][1]}
								wr.WordProperties = append(wr.WordProperties, &wp)
							}
						}
						wr.Word = request.Word
						wra = append(wra, &wr)
					}
					response := &proto.FindWordResponse{
						Result:      "OK",
						WordRecords: wra,
					}

					return response, nil
				case "async":
					cmd := common.Command{}
					cmd.Cmd = "find"
					cmd.Word = request.Word

					cmdCh <- &cmd

					ca, err := execCmd()
					result := ""
					errorValue := ""
					if err != nil {
						result = "Error"
						errorValue = fmt.Sprintf("%v", err)
					} else {
						result = ca.Result
						errorValue = ca.Error
					}

					response := &proto.FindWordResponse{
						Result: result,
						Error:  &errorValue,
						ReqId:  &ca.ID,
					}

					return response, nil

				}
				return nil, errors.New("bad mode")
			},
			checkFind: func(c context.Context, request *proto.CheckFindRequest) (*proto.CheckFindResponse, error) {
				//fmt.Printf("%v\r\n", request.ReqId)
				switch s.Mode {
				case "sync":

				case "async":
					cmd := common.Command{}
					cmd.Cmd = "check_find"
					cmd.ID = request.ReqId

					cmdCh <- &cmd

					ca, err := execCmd()
					result := "OK"
					var errorValue *string
					var wra []*proto.WordRecord

					if err != nil {
						result = "Error"
						e := fmt.Sprintf("%v", err)
						errorValue = &e
					} else {
						result = ca.Result
						errorValue = &ca.Error
						if result == "OK" {
							wra = []*proto.WordRecord{}
							for i := range ca.Words {
								wr := proto.WordRecord{}
								for j := range ca.Words[i] {
									if len(ca.Words[i][j][1]) > 0 {
										wp := proto.WordProperty{Property: ca.Words[i][j][0], Value: ca.Words[i][j][1]}
										wr.WordProperties = append(wr.WordProperties, &wp)
									}
								}
								wr.Word = ca.Word
								wra = append(wra, &wr)
							}
						}
					}

					response := &proto.CheckFindResponse{
						Result:      result,
						Error:       errorValue,
						WordRecords: wra,
					}
					return response, nil
				}
				return nil, errors.New("bad mode")
			},
			stat: func(ctx context.Context, request *proto.StatRequest) (*proto.StatResponse, error) {
				result := ""
				response := &proto.StatResponse{
					Result: result,
				}
				return response, nil
			},
		*/
	}
	serverItem.s = s
	serverItem.oc = oc
	serverItem.execCmd = execCmd
	serverItem.cmdCh = cmdCh
	serverItem.answerCh = answerCh

	proto.RegisterOpCorporaServiceServer(grpcServer, &serverItem)
	grpcServer.Serve(listener)

	oc.CloseSmallDB()
	return nil
}
