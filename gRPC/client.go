package gRPC

import (
	"context"
	"fmt"

	"github.com/wanderer69/OpCorpora/gRPC/proto"

	. "github.com/wanderer69/OpCorpora/common"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func G_RPC_init(port int) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	ss := fmt.Sprintf("127.0.0.1:%v", port) // 5300
	conn, err := grpc.Dial(ss, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
		return nil, err
	}
	return conn, nil
}

func G_RPC_Check(conn *grpc.ClientConn, query string) (string, error) {
	client := proto.NewOpCorporaServiceClient(conn)
	request := &proto.CheckRequest{
		Query: query,
	}
	response, err := client.Check(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
		return "error", err
	}
	//fmt.Println(response)
	return response.Result, nil
}
/*
type WordItem struct {
	BaseWord string
	Properties map[string]string
}
*/
func G_RPC_FindWord(conn *grpc.ClientConn, word string) (string, []WordItem, error) {
	client := proto.NewOpCorporaServiceClient(conn)
	request := &proto.FindWordRequest{
		Word: word,
	}
	//result := []map[string]string{}
	response, err := client.FindWord(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
		return "", nil, err
	}
	//fmt.Println(response)
        wi := []WordItem{}
	for i, _ := range response.WordRecords {
		vv := make(map[string]string)
		wii := WordItem{}
		wii.BaseWord = response.WordRecords[i].Baseform  
		for j, _ := range response.WordRecords[i].WordProperties {
			vv[response.WordRecords[i].WordProperties[j].Property] = response.WordRecords[i].WordProperties[j].Value
		}
                wii.Properties = vv
		//result = append(result, vv)
		wi = append(wi, wii)
	}

	return response.Result, wi, nil
}
