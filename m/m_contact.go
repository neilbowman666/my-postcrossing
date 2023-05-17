package m

import (
	"gorm.io/gorm"
	"my-postcrossing/const"
)

type Contact struct {
	gorm.Model  `json:"-"`
	Name        string
	ZipCode     string
	District    string
	Address     string
	PhoneNumber string
	Type        _const.RecipientType
}

func (_this Contact) ToVo() Visible {
	return ContactVo{
		ID:          _this.Model.ID,
		Name:        _this.Name,
		ZipCode:     _this.ZipCode,
		District:    _this.District,
		Address:     _this.Address,
		PhoneNumber: _this.PhoneNumber,
		Type:        _this.Type,
	}
}

type ContactVo struct {
	ID          uint                 `json:"id"`
	Name        string               `json:"name"`
	ZipCode     string               `json:"zip_code"`
	District    string               `json:"district"`
	Address     string               `json:"address"`
	PhoneNumber string               `json:"phone_number"`
	Type        _const.RecipientType `json:"type"`
}

func (_this ContactVo) Visible() {}
