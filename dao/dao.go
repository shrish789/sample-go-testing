package dao

import (
	"gorm.io/gorm"
)

type DataLayer interface {
	Example() bool
}

type Dao struct {
	ORM *gorm.DB
}

func GetDbInstance() (*gorm.DB) {
	return &gorm.DB{}
}

func InitializeDao() DataLayer {
	return &Dao{
		ORM: GetDbInstance(),
	}
}