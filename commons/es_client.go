package commons

import (
	"context"
	"encoding/json"
	"errors"
	"es-client/models"
	"fmt"
	"log"
	"strings"

	elasticv8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/olivere/elastic/v6"
	elasticv7 "github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
)

var (
	config        models.Config
	CurrentConfig *models.ESConfig

	clientV6 *elastic.Client
	clientV7 *elasticv7.Client
	clientV8 *elasticv8.Client
)

func init() {
	// Initialize viper to read the yaml configuration file
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			config = models.Config{}
		} else {
			// Config file was found but another error was produced
			log.Fatalf("Error reading config file: %s", err)
		}
	} else {
		if err := viper.Unmarshal(&config); err != nil {
			log.Fatalf("Unable to decode into struct: %s", err)
		}

		// Set currentConfig if a config is selected
		for i := range config.ES.Conf {
			if config.ES.Conf[i].Selected {
				CurrentConfig = &config.ES.Conf[i]
				break
			}
		}
	}
}

func InitESClient() (models.Config, error) {
	var err error
	switch CurrentConfig.Version {
	case "6":
		clientV6, err = elastic.NewClient(
			elastic.SetURL(CurrentConfig.Address),
			elastic.SetBasicAuth(CurrentConfig.Username, CurrentConfig.Password))
	case "7":
		clientV7, err = elasticv7.NewClient(
			elasticv7.SetURL(CurrentConfig.Address),
			elasticv7.SetBasicAuth(CurrentConfig.Username, CurrentConfig.Password))
	case "8":
		clientV8, err = elasticv8.NewClient(elasticv8.Config{Addresses: []string{CurrentConfig.Address}})
	default:
		err = errors.New("不支持的ES版本：" + CurrentConfig.Version)
	}
	return config, err
}

func saveConfigs() error {
	viper.Set("app.port", config.App.Port)
	viper.Set("es.conf", config.ES.Conf)
	if err := viper.WriteConfig(); err != nil {
		return err
	}
	return nil
}

func GetESConfigs() []models.ESConfig {
	return config.ES.Conf
}

func AddESConfig(newConfig models.ESConfig) string {
	for _, conf := range config.ES.Conf {
		if conf.Name == newConfig.Name {
			return "配置名已存在"
		}
	}
	config.ES.Conf = append(config.ES.Conf, newConfig)
	err := saveConfigs()
	if err != nil {
		return err.Error()
	}
	return ""
}

func DeleteESConfig(name string) string {
	for i, cfg := range config.ES.Conf {
		if cfg.Name == name {
			if CurrentConfig != nil && CurrentConfig.Name == name {
				return "已启用配置不能删除"
			}
			config.ES.Conf = append(config.ES.Conf[:i], config.ES.Conf[i+1:]...)
			err := saveConfigs()
			if err != nil {
				return err.Error()
			}
			return ""
		}
	}
	return "配置名未找到"
}

func SelectESConfig(newConfig models.ESConfig) (interface{}, error) {
	for i, cfg := range config.ES.Conf {
		if cfg.Name == newConfig.Name {
			config.ES.Conf[i].Selected = true
			CurrentConfig = &config.ES.Conf[i]
			_, err := InitESClient()
			if err != nil {
				config.ES.Conf[i].Selected = false
				CurrentConfig = nil
				return nil, err
			}
		} else {
			config.ES.Conf[i].Selected = false
		}
	}
	err := saveConfigs()
	if err != nil {
		return nil, err
	}
	return CurrentConfig, nil
}

func GetIndices() ([]string, error) {
	if CurrentConfig == nil {
		return nil, errors.New("未启用ES配置")
	}

	var indices []string
	var err error

	switch CurrentConfig.Version {
	case "6":
		indices, err = getIndicesV6()
	case "7":
		indices, err = getIndicesV7()
	case "8":
		indices, err = getIndicesV8()
	default:
		return nil, errors.New("不支持的ES版本：" + CurrentConfig.Version)
	}
	var result []string
	for _, s := range indices {
		if !strings.HasPrefix(s, ".") {
			result = append(result, s)
		}
	}
	return result, err
}

func getIndicesV6() ([]string, error) {
	res, err := clientV6.IndexNames()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func getIndicesV7() ([]string, error) {
	res, err := clientV7.IndexNames()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func getIndicesV8() ([]string, error) {
	res, err := clientV8.Cat.Indices(clientV8.Cat.Indices.WithFormat("json"))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var indices []map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&indices); err != nil {
		return nil, err
	}

	var names []string
	for _, index := range indices {
		if name, ok := index["index"].(string); ok {
			names = append(names, name)
		}
	}
	return names, nil
}

func GetMappings(index string) (map[string]interface{}, error) {
	if CurrentConfig == nil {
		return nil, errors.New("未设置ES配置")
	}

	var mappings map[string]interface{}
	var err error

	switch CurrentConfig.Version {
	case "6":
		mappings, err = getMappingsV6(index)
	case "7":
		mappings, err = getMappingsV7(index)
	case "8":
		mappings, err = getMappingsV8(index)
	default:
		return nil, errors.New("不支持的ES版本：" + CurrentConfig.Version)
	}

	return mappings, err
}

func getMappingsV6(index string) (map[string]interface{}, error) {
	res, err := clientV6.GetMapping().Index(index).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func getMappingsV7(index string) (map[string]interface{}, error) {
	res, err := clientV7.GetMapping().Index(index).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func getMappingsV8(index string) (map[string]interface{}, error) {
	res, err := clientV8.Indices.GetMapping(clientV8.Indices.GetMapping.WithIndex(index))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var mappings map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&mappings); err != nil {
		return nil, err
	}

	return mappings, nil
}

func QueryES(requestBody *models.EsSearch) (*models.EsData, error) {
	if CurrentConfig == nil {
		return nil, errors.New("未设置ES配置")
	}
	var res *models.EsData
	var err error

	switch CurrentConfig.Version {
	case "6":
		res, err = queryESV6(requestBody)
	case "7":
		res, err = queryESV7(requestBody)
	case "8":
		res, err = queryESV8(requestBody)
	default:
		return nil, errors.New("不支持的ES版本：" + CurrentConfig.Version)
	}

	if err != nil {
		return nil, err
	}
	return res, nil
}

func queryESV6(requestBody *models.EsSearch) (*models.EsData, error) {
	from := (requestBody.CurrentPage - 1) * requestBody.PageSize
	if from > 1000 {
		from = 1000
	}
	requestBody.QueryJson["from"] = from
	requestBody.QueryJson["size"] = requestBody.PageSize
	searchResult, err := clientV6.Search().
		Index(requestBody.Index).
		Source(requestBody.QueryJson).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	response := models.EsData{
		Total: int(searchResult.TotalHits()),
		Data:  make([]interface{}, 0),
	}

	for _, hit := range searchResult.Hits.Hits {
		var doc map[string]interface{}
		if err := json.Unmarshal(*hit.Source, &doc); err != nil {
			return nil, err
		} else {
			response.Data = append(response.Data, doc)
		}
	}

	return &response, nil
}

func queryESV7(requestBody *models.EsSearch) (*models.EsData, error) {
	from := (requestBody.CurrentPage - 1) * requestBody.PageSize
	if from > 1000 {
		from = 1000
	}
	requestBody.QueryJson["from"] = from
	requestBody.QueryJson["size"] = requestBody.PageSize
	searchResult, err := clientV7.Search().
		Index(requestBody.Index).
		Source(requestBody.QueryJson).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	response := models.EsData{
		Total: int(searchResult.TotalHits()),
		Data:  make([]interface{}, 0),
	}

	for _, hit := range searchResult.Hits.Hits {
		var doc map[string]interface{}
		if err := json.Unmarshal(hit.Source, &doc); err != nil {
			return nil, err
		} else {
			response.Data = append(response.Data, doc)
		}
	}

	return &response, nil
}

func queryESV8(requestBody *models.EsSearch) (*models.EsData, error) {
	from := (requestBody.CurrentPage - 1) * requestBody.PageSize
	if from > 1000 {
		from = 1000
	}
	searchBody := fmt.Sprintf(`{
        "from": %d,
        "size": %d,
        %s
    }`, from, requestBody.PageSize, requestBody.QueryJson)

	// 执行查询
	res, err := clientV8.Search(
		clientV8.Search.WithContext(context.Background()),
		clientV8.Search.WithIndex(requestBody.Index),
		clientV8.Search.WithBody(strings.NewReader(searchBody)),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// 处理查询结果
	var searchResult struct {
		Hits struct {
			Total struct {
				Value int `json:"value"`
			} `json:"total"`
			Hits []struct {
				Source json.RawMessage `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(res.Body).Decode(&searchResult); err != nil {
		return nil, err
	}

	response := models.EsData{
		Total: searchResult.Hits.Total.Value,
		Data:  make([]interface{}, 0),
	}

	for _, hit := range searchResult.Hits.Hits {
		var doc map[string]interface{}
		if err := json.Unmarshal(hit.Source, &doc); err != nil {
			return nil, err
		} else {
			response.Data = append(response.Data, doc)
		}
	}

	return &response, nil
}
