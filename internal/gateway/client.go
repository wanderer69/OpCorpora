package gateway

import (
	"context"
	"fmt"

	"github.com/wanderer69/OpCorpora/pkg/proto"

	"github.com/wanderer69/OpCorpora/public/common"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

func GRPCInit(address string, port int) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		//grpc.WithInsecure(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	ss := fmt.Sprintf("%v:%v", address, port) // 5300
	conn, err := grpc.Dial(ss, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
		return nil, err
	}
	return conn, nil
}

func GRPCInitFull(ip_addr string, port int) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		//grpc.WithInsecure(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	ss := fmt.Sprintf("%v:%v", ip_addr, port) // 5300
	conn, err := grpc.Dial(ss, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
		return nil, err
	}
	return conn, nil
}

func GRPCCheck(conn *grpc.ClientConn, query string) (string, error) {
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

func GRPCMode(conn *grpc.ClientConn, mode string) (string, error) {
	client := proto.NewOpCorporaServiceClient(conn)
	request := &proto.ModeRequest{
		Mode: mode,
	}
	response, err := client.Mode(context.Background(), request)
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
func GRPCFindWord(conn *grpc.ClientConn, word string) (string, []common.WordItem, error) {
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
	wi := []common.WordItem{}
	for i := range response.WordRecords {
		vv := make(map[string]string)
		wii := common.WordItem{}
		wii.BaseWord = response.WordRecords[i].Baseform
		for j := range response.WordRecords[i].WordProperties {
			vv[response.WordRecords[i].WordProperties[j].Property] = response.WordRecords[i].WordProperties[j].Value
		}
		wii.Properties = vv
		//result = append(result, vv)
		wi = append(wi, wii)
	}

	return response.Result, wi, nil
}
