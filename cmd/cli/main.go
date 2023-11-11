package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wanderer69/OpCorpora/internal/gateway"
)

func LoadValue(fn string) (string, error) {
	data, err := os.ReadFile(fn)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func main() {
	modePtr := flag.String("mode", "", "mode: check, findword")
	//	port := 5300

	var port int = 5300
	flag.IntVar(&port, "port", 0, "service port")

	var value string
	flag.StringVar(&value, "value", "", "value")
	var address string
	flag.StringVar(&address, "address", "", "address")
	var file_name string
	flag.StringVar(&file_name, "file_name", "", "file name")

	flag.Parse()
	switch *modePtr {
	case "check":
		cc, err := gateway.GRPCInit(address, port)
		if err != nil {
			fmt.Printf("%v\r\n", err)
		}
		res, err := gateway.GRPCCheck(cc, value)
		if err != nil {
			fmt.Printf("%v\r\n", err)
		}
		fmt.Printf("%v\r\n", res)
	case "mode":
		cc, err := gateway.GRPCInit(address, port)
		if err != nil {
			fmt.Printf("%v\r\n", err)
		}
		res, err := gateway.GRPCMode(cc, value)
		if err != nil {
			fmt.Printf("%v\r\n", err)
		}
		fmt.Printf("%v\r\n", res)
	case "findword":
		cc, err := gateway.GRPCInit(address, port)
		if err != nil {
			fmt.Printf("%v\r\n", err)
		}

		value, err = LoadValue(file_name)
		if err != nil {
			fmt.Printf("%v\r\n", err)
			return
		}

		word, bw, err := gateway.GRPCFindWord(cc, value)
		if err != nil {
			fmt.Printf("%v\r\n", err)
		}
		fmt.Printf("%v %v\r\n", word, bw)
	}
}
