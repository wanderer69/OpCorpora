package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"os"
	"os/signal"

	"github.com/wanderer69/OpCorpora/internal/gateway"
	"github.com/wanderer69/OpCorpora/internal/settings"
)

func LoadSettings() (*settings.Settings, error) {
	data, err := os.ReadFile("settings.json")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	// json data
	var s settings.Settings

	// unmarshall it
	err = json.Unmarshal(data, &s)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	return &s, nil
}

func SaveSettings(s *settings.Settings) error {
	data_1, err2_ := json.MarshalIndent(&s, "", "  ")
	if err2_ != nil {
		fmt.Println("error:", err2_)
		return err2_
	}
	_ = os.WriteFile("settings.json", data_1, 0644)
	return nil
}

func main() {
	// modePtr := flag.String("mode", "login", "mode: register, login, method")

	var port_var string
	flag.StringVar(&port_var, "port", "9091", "external port")
	var port_client_var string
	flag.StringVar(&port_client_var, "port_client", "5300", "gRPC port")

	flag.Parse()

	s, err_ := LoadSettings()
	if err_ != nil {
		fmt.Println(err_)
		s := settings.Settings{}
		s.PortClient = 5300
		s.OpCorporaPath = "C:/Development/Go projects/OpCorpora/OCorporaDBFull"
		s.Mode = "async"
		err := SaveSettings(&s)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	// MyPageServer("/api/v1", s)    // /newpage
	fmt.Printf("Server started\r\n")
	gateway.GRPCServer(s)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit
}
