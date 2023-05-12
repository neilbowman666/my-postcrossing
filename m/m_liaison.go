package m

import (
	"gorm.io/gorm"
	"my-postcrossing/const"
)

type Liaison struct {
	gorm.Model  `json:"-"`
	Name        string
	ZipCode     string
	District    string
	Address     string
	PhoneNumber string
	Type        _const.RecipientType
}
