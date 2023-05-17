package form

import (
	_const "my-postcrossing/const"
	"my-postcrossing/m"
)

type AddContactForm struct {
	Name        string               `json:"name" example:"Leon Kennedy"`
	ZipCode     string               `json:"zip_code" example:"100177"`
	District    string               `json:"district" example:"ChengduShi QingyangQu"`
	Address     string               `json:"address" example:"CaotangLu 1 Hao"`
	PhoneNumber string               `json:"phone_number" example:"028-87009999"`
	Type        _const.RecipientType `json:"type" example:"friend"`
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
