package m

import (
	"gorm.io/gorm"
	"my-postcrossing/const"
	"time"
)

type SentMail struct {
	gorm.Model `json:"-"`

	SendDate time.Time

	MailType _const.MailType

	PostFee            float32
	PostFeePaymentType _const.PostFeePaymentType

	PostRegion _const.PostRegion

	RecipientType        _const.RecipientType
	RecipientName        string
	RecipientZipCode     string
	RecipientDistrict    string
	RecipientAddress     string
	RecipientPhoneNumber string
}

func (_this SentMail) ToVo() Visible {
	return SentMailVo{
		ID:                   _this.Model.ID,
		SendDate:             _this.SendDate,
		MailType:             _this.MailType,
		PostFee:              _this.PostFee,
		PostFeePaymentType:   _this.PostFeePaymentType,
		PostRegion:           _this.PostRegion,
		RecipientType:        _this.RecipientType,
		RecipientName:        _this.RecipientName,
		RecipientZipCode:     _this.RecipientZipCode,
		RecipientDistrict:    _this.RecipientDistrict,
		RecipientAddress:     _this.RecipientAddress,
		RecipientPhoneNumber: _this.RecipientPhoneNumber,
	}
}

type SentMailVo struct {
	ID uint `json:"id"`

	SendDate time.Time `json:"send_date"`

	MailType _const.MailType `json:"mail_type"`

	PostFee            float32                   `json:"post_fee"`
	PostFeePaymentType _const.PostFeePaymentType `json:"post_fee_payment_type"`

	PostRegion _const.PostRegion `json:"post_region"`

	RecipientType        _const.RecipientType `json:"recipient_type"`
	RecipientName        string               `json:"recipient_name"`
	RecipientZipCode     string               `json:"recipient_zip_code"`
	RecipientDistrict    string               `json:"recipient_district"`
	RecipientAddress     string               `json:"recipient_address"`
	RecipientPhoneNumber string               `json:"recipient_phone_number"`
}

func (SentMailVo) Visible() {

}
