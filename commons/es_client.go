package commons

import (
	"context"
	"encoding/json"
	"es-client/models"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

var es *elasticsearch.Client

// 设置es客户端
func InitESClient(esConfig *models.EsConfig) {
	cfg := elasticsearch.Config{
		Addresses: esConfig.Addresses,
		Username:  esConfig.Username,
		Password:  esConfig.Password,
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Printf("Error creating the client: %s", err)
	}
	es = client
}

// 解析 Elasticsearch 响应
func decodeResponse(res *esapi.Response, target interface{}) error {
	if res.IsError() {
		return fmt.Errorf("error response: %s", res.String())
	}
	err := json.NewDecoder(res.Body).Decode(target)
	if err != nil {
		return err
	}
	return nil
}

// 获取es连接信息
func CheckESClient() map[string]interface{} {
	res, err := es.Info()
	if err != nil {
		log.Printf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	// 检查响应状态
	if res.IsError() {
		log.Printf("Error response: %s", res.String())
	}

	// 解析响应体
	var responseBody map[string]interface{}
	if err := decodeResponse(res, &responseBody); err != nil {
		log.Printf("Error parsing response body: %s", err)
	}

	return responseBody
}

// 获取所有索引
func GetIndices() []string {
	// 执行获取所有索引的操作
	res, err := esapi.CatIndicesRequest{Format: "json"}.Do(context.Background(), es)

	if err != nil {
		log.Printf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	// 检查响应状态
	if res.IsError() {
		log.Printf("Error response: %s", res.String())
	}

	// 解析响应体
	var responseBody []map[string]interface{}
	if err := decodeResponse(res, &responseBody); err != nil {
		log.Printf("Error parsing response body: %s", err)
	}

	// 提取索引名称
	var indexNames []string
	for _, row := range responseBody {
		indexName := row["index"].(string)
		indexNames = append(indexNames, indexName)
	}

	return indexNames
}

// 更加索引名称获取索引字段映射
func GetIndexMapping(indexNames []string) map[string]interface{} {
	// 创建 IndicesGetFieldMapping 请求
	req := esapi.IndicesGetMappingRequest{
		Index: indexNames, // 设置要获取字段映射的索引
	}

	res, err := req.Do(context.Background(), es)

	if err != nil {
		log.Printf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	// 检查响应状态
	if res.IsError() {
		log.Printf("Error response: %s", res.String())
	}

	// 解析响应体
	var responseBody map[string]interface{}
	if err := decodeResponse(res, &responseBody); err != nil {
		log.Printf("Error parsing response body: %s", err)
	}

	return responseBody
}
