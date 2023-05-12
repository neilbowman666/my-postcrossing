package m

import "gorm.io/gorm"

type Liaison struct {
	gorm.Model  `json:"-"`
	Name        string
	ZipCode     string
	District    string
	Address     string
	PhoneNumber string
}
