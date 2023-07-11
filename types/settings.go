package types

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Settings struct {
	SelectedProducts []string `json:"selected_products"`
}

// Checks if an item has been selected
func (s *Settings) ItemSelected(c string) bool {
	for _, v := range s.SelectedProducts {
		if v == c {
			return true
		}
	}

	return false
}

// Save settings to a json
func (s Settings) SaveSettings() error {
    if _, err := os.Stat("data/config.json"); os.IsNotExist(err) {
        os.Mkdir("data", 0755)
    }

	json, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("data/config.json", json, 0755)
	if err != nil {
		return err
	}

	return nil
}

// Loads settings from a json
func (s *Settings) LoadSettings() error {
    data, err := ioutil.ReadFile("data/config.json")
    if err != nil {
        return err
    }
    
    tempSettings := Settings{
        SelectedProducts: make([]string, 0),
    }

    err = json.Unmarshal([]byte(data), &tempSettings)
    if err != nil {
        return err
    }

    *s = tempSettings

    return nil
}
