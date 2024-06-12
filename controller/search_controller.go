package controller

import (
	"es-client/commons"
	"es-client/models"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
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
func (con SearchController) GetIndices(c *gin.Context) {
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
func (con SearchController) GetMappings(c *gin.Context) {
	index := c.Query("index")
	mappings, err := commons.GetMappings(index)

	if err != nil {
		con.Error(c, "获取索引mappings失败", err.Error())
		return
	}

	result := SortMappings(mappings, index)
	con.Ok(c, "获取索引mappings成功", result)
}

func SortMappings(mappings map[string]interface{}, index string) []map[string]interface{} {
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
					"value":    k1,
					"label":    k1,
					"type":     v1["type"],
					"children": make([]interface{}, 0),
				})
			}
		} else {
			pNode["children"] = make([]interface{}, 0)
		}
		result = append(result, pNode)
	}
	return result
}

// QueryES
// @Summary 查询es
// @Tags es查询
// @Accept json
// @Produce json
// @Param        queryData    body     models.EsSearch  false  "es查询"
// @Success 200 {string} json{"code","msg","data"}
// @Router /ec/es/queryES [post]
func (con SearchController) QueryES(c *gin.Context) {
	var queryData models.EsSearch
	if err := c.ShouldBindJSON(&queryData); err != nil {
		con.Error(c, "查询失败", err.Error())
		return
	}
	res, err := commons.QueryES(&queryData)
	if err != nil {
		con.Error(c, "查询失败", err.Error())
		return
	}
	con.Ok(c, "查询成功", res)
}

// ExportES
// @Summary 导出es
// @Tags es查询
// @Accept json
// @Produce json
// @Param        queryData    body     models.EsSearch  false  "es查询"
// @Success 200 {string} json{"code","msg","data"}
// @Router /ec/es/exportES [post]
func (con SearchController) ExportES(c *gin.Context) {
	var queryData models.EsSearch
	if err := c.ShouldBindJSON(&queryData); err != nil {
		con.Error(c, "导出失败", err.Error())
		return
	}
	queryData.CurrentPage = 1
	queryData.PageSize = 10000
	res, err := commons.QueryES(&queryData)
	if err != nil {
		con.Error(c, "导出失败", err.Error())
		return
	}
	f := excelize.NewFile()
	sheet := "result"
	// f.NewSheet(sheet)
	f.SetSheetName("Sheet1", sheet)
	mappings, _ := commons.GetMappings(queryData.Index)
	headers := SortMappings(mappings, queryData.Index)
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, header)
	}

	for i, hit := range res.Data {
		hitMap := hit.(map[string]interface{})
		for j, header := range headers {
			cell, _ := excelize.CoordinatesToCellName(j+1, i+2)
			f.SetCellValue(sheet, cell, hitMap[header["value"].(string)])
		}
	}
	dirPath := "./exportData/"
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// 如果文件夹不存在，则创建文件夹
		err := os.Mkdir(dirPath, os.ModePerm)
		con.Error(c, "导出文件夹创建失败", err.Error())
		return
	}
	timestampStr := strconv.FormatInt(time.Now().Unix(), 10)
	filename := dirPath + timestampStr + ".xlsx"
	if err := f.SaveAs(filename); err != nil {
		con.Error(c, "导出失败", err.Error())
		return
	}

	// buffer, err := f.WriteToBuffer()
	// if err != nil {
	// 	con.Error(c, "导出失败", err.Error())
	// 	return
	// }

	// timestampStr := strconv.FormatInt(time.Now().Unix(), 10)
	// filename := timestampStr + ".xlsx"
	// // c.Writer.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	// c.Writer.Header().Add("Content-Type", "application/octet-stream")
	// c.Writer.Header().Add("Content-Disposition", "attachment; filename=" + filename)
	// // c.Writer.Header().Add("Content-Length", strconv.Itoa(buffer.Len()))
	// c.Writer.Write(buffer.Bytes())
	con.Ok(c, "导出成功，文件路径："+filename, nil)
}
