package models

import "gorm.io/gorm"

type EsConfig struct {
	gorm.Model
	Name     string
	Host     string
	UserName string
	Password string
}

func (ec *EsConfig) TableName() string {
	return "es_config"
}
