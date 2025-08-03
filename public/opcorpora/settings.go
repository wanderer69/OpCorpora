package opcorpora

import (
	"encoding/json"
	"fmt"
	"os"
)

type Settings struct {
	RestrictionsList []string `json:"restrictions_list"`
}

func InitSettings() *Settings {
	settings := Settings{}
	settings.RestrictionsList = []string{"Abbr", "Name", "Surn", "Patr", "Geox", "Orgn", "Trad", "Init", "Hypo"}
	return &settings
}

func LoadSettings(file_name string) (*Settings, error) {
	data, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	var settings Settings
	err = json.Unmarshal(data, &settings)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	return &settings, nil
}

func SaveSettings(file_name string, settings *Settings) error {
	data, err := json.Marshal(&settings)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	_ = os.WriteFile(file_name, data, 0644)
	return nil
}
