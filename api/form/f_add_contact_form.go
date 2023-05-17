package form

import (
	_const "my-postcrossing/const"
	"my-postcrossing/m"
)

type AddContactForm struct {
	Name        string               `json:"name"`
	ZipCode     string               `json:"zip_code"`
	District    string               `json:"district"`
	Address     string               `json:"address"`
	PhoneNumber string               `json:"phone_number"`
	Type        _const.RecipientType `json:"type"`
}

func (f AddContactForm) ToM() *m.Contact {
	return &m.Contact{
		Name:        f.Name,
		ZipCode:     f.ZipCode,
		District:    f.District,
		Address:     f.Address,
		PhoneNumber: f.PhoneNumber,
		Type:        f.Type,
	}
}
