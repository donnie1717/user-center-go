package mysql

import "github.com/jinzhu/gorm"

type datastore struct {
	db *gorm.DB
}

type mysql struct {
	dsn string `yaml:"dsn"`
}