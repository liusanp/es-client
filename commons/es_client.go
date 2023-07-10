package commons

import (
	"es-client/models"
	"log"

	ES7 "github.com/elastic/go-elasticsearch/v7"
	ES8 "github.com/elastic/go-elasticsearch/v8"
)

var ES7Client *ES7.Client
var ES8Client *ES8.Client

func InitESClient(esConfig *models.EsConfig) {
	if esConfig.Version == 6 || esConfig.Version == 7 {
		cfg := ES7.Config{
			Addresses: esConfig.Addresses,
			Username:  esConfig.Username,
			Password:  esConfig.Password,
		}
		es, err := ES7.NewClient(cfg)
		if err != nil {
			log.Fatalf("Error creating the client: %s", err)
		}
		ES7Client = es
	}
	if esConfig.Version == 8 {
		cfg := ES8.Config{
			Addresses: esConfig.Addresses,
			Username:  esConfig.Username,
			Password:  esConfig.Password,
		}
		es, err := ES8.NewClient(cfg)
		if err != nil {
			log.Fatalf("Error creating the client: %s", err)
		}
		ES8Client = es
	}
}

func CheckESClient(version uint) {
	if version == 6 || version == 7 {
		res, err := ES7Client.Info()
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		defer res.Body.Close()
		log.Println(res)
	}
	if version == 8 {
		res, err := ES8Client.Info()
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		defer res.Body.Close()
		log.Println(res)
	}
}
