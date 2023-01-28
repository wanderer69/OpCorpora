package gRPC

import (
    "net"
    "github.com/wanderer69/OpCorpora/gRPC/proto"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/grpclog"
    "fmt"
    "errors"
	. "github.com/wanderer69/OpCorpora/settings"
	moc "github.com/wanderer69/OpCorpora"    
	. "github.com/wanderer69/OpCorpora/common"
	proc "github.com/wanderer69/OpCorpora/process"
)

type server struct{
     mode func (context.Context, *proto.ModeRequest) (*proto.ModeResponse, error)
     check func (context.Context, *proto.CheckRequest) (*proto.CheckResponse, error)
     find_word func (context.Context, *proto.FindWordRequest) (*proto.FindWordResponse, error)
     check_find func (context.Context, *proto.CheckFindRequest) (*proto.CheckFindResponse, error)
     stat func (context.Context, *proto.StatRequest) (*proto.StatResponse, error)
}

func (s *server) Mode (ctx context.Context, in *proto.ModeRequest) (*proto.ModeResponse, error) {
    r, err := s.mode(ctx, in)
    return r, err
}

func (s *server) Check (ctx context.Context, in *proto.CheckRequest) (*proto.CheckResponse, error) {
    r, err := s.check(ctx, in)
    return r, err
}

func (s *server) FindWord (ctx context.Context, in *proto.FindWordRequest) (*proto.FindWordResponse, error) {
    r, err := s.find_word(ctx, in)
    return r, err
}

func (s *server) CheckFind (ctx context.Context, in *proto.CheckFindRequest) (*proto.CheckFindResponse, error) {
    r, err := s.check_find(ctx, in)
    return r, err
}

func (s *server) Stat (ctx context.Context, in *proto.StatRequest) (*proto.StatResponse, error) {
    r, err := s.stat(ctx, in)
    return r, err
}

type ModeFunc func (ctx context.Context, mode string) (string, error)
type CheckFunc func (ctx context.Context, query string) (string, error)
type FindWordFunc func (ctx context.Context, word string) (string, string, string, map[string]string, error)
type CheckFindFunc func (ctx context.Context, id string) (string, string, string, map[string]string, error)
type StatFunc func (ctx context.Context, mode string) (string, string, string, map[string]string, error)

func G_RPC_server(s *Settings/*port int, path string, flag_mode string*/) error {
	ss := fmt.Sprintf(":%v", s.PortClient) // 5300
    listener, err := net.Listen("tcp", ss)

    if err != nil {
        grpclog.Fatalf("failed to listen: %v", err)
        return err
    }

	var cmd_ch chan *Command
	var answer_ch chan *CommandAnswer
	//var oc *moc.OCorpora

	oc, err := moc.OpenOCorporaFull(s.OpCorporaPath, 0)
	if err != nil {
		grpclog.Fatalf("failed to open db: %v", err)
		return err		
	}
	
	init_async := func(oc *moc.OCorpora) {
		cmd_ch = make(chan *Command)
		answer_ch = make(chan *CommandAnswer)

		go proc.Tasker(s, cmd_ch, answer_ch, oc)
	}
	exec_cmd := func(ca *CommandAnswer) (*CommandAnswer, error) {
		// надо дождаться ответа
		ca = nil
		flag := false
		for {
			select {
					case ca = <- answer_ch:
						flag = true
			}
			if flag {
				break
			}
		}
		return ca, nil
	}

	switch s.Mode {
		case "sync":

		case "async":
			init_async(oc)
	}
	
    opts := []grpc.ServerOption{}
    grpcServer := grpc.NewServer(opts...)
    server_item := server {  
        check : func(c context.Context, request *proto.CheckRequest) (*proto.CheckResponse, error) {
            // fmt.Printf("%v\r\n", request.Query)
            switch s.Mode {
				case "sync":
				case "async":
					cmd := Command{}
					cmd.Cmd = "check"
					cmd.ID = request.Query
					cmd_ch <- &cmd	
					var ca *CommandAnswer
					ca, err := exec_cmd(ca)
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
			return nil, errors.New("Bad mode") 
        },
        mode : func(c context.Context, request *proto.ModeRequest) (*proto.ModeResponse, error) {
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
				Result   : result,
            }
            return response, nil
        },
        find_word : func(c context.Context, request *proto.FindWordRequest) (*proto.FindWordResponse, error) {
            //fmt.Printf("%v\r\n", request.Token)
            switch s.Mode {
				case "sync":
					result, err := moc.FindWordCorporaService(oc, request.Word)
					if err != nil {
						response := &proto.FindWordResponse{
							Result : "Error",
						}
						return response, nil    
					}
					wra := []*proto.WordRecord{}
					for i, _ := range result {
						wr := proto.WordRecord{}
						for j, _ := range result[i] {
								if len(result[i][j][1]) > 0 {
								wp := proto.WordProperty{Property: result[i][j][0], Value: result[i][j][1]}
								wr.WordProperties = append(wr.WordProperties, &wp)
							}
						}
						wr.Word = request.Word
						wra = append(wra, &wr)
					}
					response := &proto.FindWordResponse{
						Result : "OK",
						WordRecords : wra,
					}
				
					return response, nil    
				case "async":
					cmd := Command{}
					cmd.Cmd = "find"
					cmd.Word = request.Word
					
					cmd_ch <- &cmd	

					var ca *CommandAnswer
					ca, err := exec_cmd(ca)
					result := ""
					error_ := ""
					if err != nil {
						result = "Error"
						error_ = fmt.Sprintf("%v", err)
					} else {
						result = ca.Result
						error_ = ca.Error
					}

					response := &proto.FindWordResponse{
						Result: result,
						Error: &error_,
						ReqId: &ca.ID,
					}
				
					return response, nil    
				
			}
			return nil, errors.New("Bad mode") 
        },
        check_find : func(c context.Context, request *proto.CheckFindRequest) (*proto.CheckFindResponse, error) {
            //fmt.Printf("%v\r\n", request.ReqId)
            switch s.Mode {
				case "sync":
		
				case "async":
					cmd := Command{}
					cmd.Cmd = "check_find"
					cmd.ID = request.ReqId
					
					cmd_ch <- &cmd	

					var ca *CommandAnswer
					ca, err := exec_cmd(ca)
					result := "OK"
					var error_ *string
					var wra []*proto.WordRecord

					if err != nil {
						result = "Error"
						e := fmt.Sprintf("%v", err)
						error_ = &e
					} else {
						result = ca.Result
						error_ = &ca.Error
						if result == "OK" {
							wra = []*proto.WordRecord{}
							for i, _ := range ca.Words {
								wr := proto.WordRecord{}
								for j, _ := range ca.Words[i] {
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
						Result: result,
						Error: error_,
						WordRecords : wra,
					}
					return response, nil
			}
			return nil, errors.New("Bad mode") 
        },
		stat: func(ctx context.Context, request *proto.StatRequest) (*proto.StatResponse, error) {
			result := ""
			response := &proto.StatResponse{
				Result: result,
			}
			return response, nil
		},
    }
    proto.RegisterOpCorporaServiceServer(grpcServer, proto.OpCorporaServiceServer(&server_item))
    grpcServer.Serve(listener)

	oc.CloseSmallDB()
    return nil
}

