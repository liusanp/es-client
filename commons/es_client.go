package commons

import (
	"context"
	"encoding/json"
	"errors"
	"es-client/models"
	"log"

	elasticv8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/olivere/elastic/v6"
	elasticv7 "github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
)

var (
	config        models.Config
	currentConfig *models.ESConfig

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
				currentConfig = &config.ES.Conf[i]
				break
			}
		}
	}
}

func InitESClient() models.Config {
	var err error
	switch currentConfig.Version {
	case "6":
		clientV6, err = elastic.NewClient(elastic.SetURL(currentConfig.Address))
	case "7":
		clientV7, err = elasticv7.NewClient(elasticv7.SetURL(currentConfig.Address))
	case "8":
		clientV8, err = elasticv8.NewClient(elasticv8.Config{Addresses: []string{currentConfig.Address}})
	default:
		log.Fatalf("不支持的ES版本：" + currentConfig.Version)
	}
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}
	return config
}

func saveConfigs() {
	viper.Set("app.port", config.App.Port)
	viper.Set("es.conf", config.ES.Conf)
	if err := viper.WriteConfig(); err != nil {
		log.Fatalf("Error writing config file: %s", err)
	}
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
	saveConfigs()
	return ""
}

func DeleteESConfig(name string) string {
	for i, cfg := range config.ES.Conf {
		if cfg.Name == name {
			config.ES.Conf = append(config.ES.Conf[:i], config.ES.Conf[i+1:]...)
			if currentConfig != nil && currentConfig.Name == name {
				currentConfig = nil
				clientV6 = nil
				clientV7 = nil
				clientV8 = nil
			}
			saveConfigs()
			return ""
		}
	}
	return "配置名未找到"
}

func SelectESConfig(newConfig models.ESConfig) interface{} {
	for i, cfg := range config.ES.Conf {
		if cfg.Name == newConfig.Name {
			config.ES.Conf[i].Selected = true
			currentConfig = &config.ES.Conf[i]
			InitESClient()
		} else {
			config.ES.Conf[i].Selected = false
		}
	}
	saveConfigs()
	return currentConfig
}

func GetIndices() ([]string, error) {
	if currentConfig == nil {
		return nil, errors.New("未设置ES配置")
	}

	var indices []string
	var err error

	switch currentConfig.Version {
	case "6":
		indices, err = getIndicesV6()
	case "7":
		indices, err = getIndicesV7()
	case "8":
		indices, err = getIndicesV8()
	default:
		return nil, errors.New("不支持的ES版本：" + currentConfig.Version)
	}

	return indices, err
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
	if currentConfig == nil {
		return nil, errors.New("未设置ES配置")
	}

	var mappings map[string]interface{}
	var err error

	switch currentConfig.Version {
	case "6":
		mappings, err = getMappingsV6(index)
	case "7":
		mappings, err = getMappingsV7(index)
	case "8":
		mappings, err = getMappingsV8(index)
	default:
		return nil, errors.New("不支持的ES版本：" + currentConfig.Version)
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
