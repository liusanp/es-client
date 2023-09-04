package commons

import (
	"context"
	"es-client/models"
	"log"

	ES7 "github.com/elastic/go-elasticsearch/v7"
	ES7API "github.com/elastic/go-elasticsearch/v7/esapi"
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
		ES8Client = nil
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
		ES7Client = nil
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

func GetIndexMapping() {
	// 指定要获取字段映射的索引
	indexName := "sq-yshj"

	// 创建 IndicesGetFieldMapping 请求
	req := ES7API.IndicesGetFieldMappingRequest{
		Index: []string{indexName}, // 设置要获取字段映射的索引
	}

	res, err := req.Do(context.Background(), ES7Client)
	
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)
}
