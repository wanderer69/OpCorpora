package main

import (
	"fmt"
	"flag"
	. "arkhangelskiy-dv.ru/OpCorpora/gRPC"
)

func main() {
	modePtr := flag.String("mode", "", "mode: check, findword")
	port := 5300
	var value string
	flag.StringVar(&value, "value", "", "value")

	flag.Parse()
	switch *modePtr {
	case "check":
		cc, err := G_RPC_init(port)
		if err != nil {
		      fmt.Printf("%v\r\n", err)
		}
		res, err := G_RPC_Check(cc, value)
		if err != nil {
		      fmt.Printf("%v\r\n", err)
		}
		fmt.Printf("%v\r\n", res)
	case "findword":
		cc, err := G_RPC_init(port)
		if err != nil {
		      fmt.Printf("%v\r\n", err)
		}
		word, bw, res, err := G_RPC_FindWord(cc, value)
		if err != nil {
		      fmt.Printf("%v\r\n", err)
		}
		fmt.Printf("%v %v %v\r\n", word, bw, res)
	}
}
