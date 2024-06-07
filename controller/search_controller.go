package controller

import (
	"es-client/commons"
	"sort"

	"github.com/gin-gonic/gin"
)

type SearchController struct {
	BaseController
}

// GetIndices
// @Summary 获取es索引
// @Tags es查询
// @Accept json
// @Produce json
// @Success 200 {string} json{"code","msg","data"}
// @Router /ec/es/getIndices [get]
func (con BaseController) GetIndices(c *gin.Context) {
	indices, err := commons.GetIndices()
	if err != nil {
		con.Error(c, "获取索引列表失败", err.Error())
		return
	}
	sort.Strings(indices)

	con.Ok(c, "获取索引列表成功", indices)
}

// GetMappings
// @Summary 获取es索引mappings
// @Tags es查询
// @Accept json
// @Produce json
// @Param        index    query     string  false  "索引名称"
// @Success 200 {string} json{"code","msg","data"}
// @Router /ec/es/getMappings [get]
func (con BaseController) GetMappings(c *gin.Context) {
	index := c.Query("index")
	mappings, err := commons.GetMappings(index)

	if err != nil {
		con.Error(c, "获取索引mappings失败", err.Error())
		return
	}
	r := mappings[index].(map[string]interface{})["mappings"].(map[string]interface{})["index"].(map[string]interface{})["properties"].(map[string]interface{})
	result := make([]map[string]interface{}, 0)
	keys := make([]string, 0)
	for k, _ := range r {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := r[k].(map[string]interface{})
		nodeType := ""
		if v["type"] != nil {
			nodeType = v["type"].(string)
		} else if v["properties"] != nil {
			nodeType = "nested"
		}
		pNode := make(map[string]interface{})
		pNode["value"] = k
		pNode["label"] = k
		pNode["type"] = nodeType
		if nodeType == "nested" {
			props := v["properties"].(map[string]interface{})
			pNode["children"] = make([]map[string]interface{}, 0)
			for k1, sv1 := range props {
				v1 := sv1.(map[string]interface{})
				pNode["children"] = append(pNode["children"].([]map[string]interface{}), map[string]interface{}{
					"value": k1,
					"label": k1,
					"type":  v1["type"],
					"children": make([]interface{}, 0),
				})
			}
		} else {
			pNode["children"] = make([]interface{}, 0)
		}
		result = append(result, pNode)
	}

	con.Ok(c, "获取索引mappings成功", result)
}
