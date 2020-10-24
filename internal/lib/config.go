package lib

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Mynor2397/virtual-parish-office/internal/models"
)

// Config del servidor
func Config() models.Config {
	var parameters models.Config
	configfile, err := os.Open("./config/config.json")
	if err != nil {
		log.Println(err.Error())
	}

	defer configfile.Close()

	var configDecoder *json.Decoder = json.NewDecoder(configfile)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = configDecoder.Decode(&parameters)
	if err != nil {
		log.Fatal(err.Error())
	}

	return parameters
}
