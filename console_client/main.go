package main

import (
	"fmt"
	"flag"
	"io/ioutil"

	. "github.com/wanderer69/OpCorpora/gRPC"
)

func LoadValue(fn string) (string, error) {
	data, err := ioutil.ReadFile(fn)
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
		cc, err := G_RPC_init(address, port)
		if err != nil {
		      fmt.Printf("%v\r\n", err)
		}
		res, err := G_RPC_Check(cc, value)
		if err != nil {
		      fmt.Printf("%v\r\n", err)
		}
		fmt.Printf("%v\r\n", res)
	case "mode":
		cc, err := G_RPC_init(address, port)
		if err != nil {
		      fmt.Printf("%v\r\n", err)
		}
		res, err := G_RPC_Mode(cc, value)
		if err != nil {
		      fmt.Printf("%v\r\n", err)
		}
		fmt.Printf("%v\r\n", res)
	case "findword":
		cc, err := G_RPC_init(address, port)
		if err != nil {
		      fmt.Printf("%v\r\n", err)
		}

		value, err = LoadValue(file_name)
		if err != nil {
			fmt.Printf("%v\r\n", err)
			return
		}

		word, bw, err := G_RPC_FindWord(cc, value)
		if err != nil {
		      fmt.Printf("%v\r\n", err)
		}
		fmt.Printf("%v %v\r\n", word, bw)
	}
}
