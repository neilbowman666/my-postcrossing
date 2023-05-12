package db

import "gorm.io/gorm"

type Transaction struct {
	Tx *gorm.DB
}

func (tx *Transaction) Save(value interface{}) error {
	return DBConn.Save(value).Error
}
