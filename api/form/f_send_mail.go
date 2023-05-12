package form

import (
	_const "my-postcrossing/const"
	"my-postcrossing/m"
)

type SendMailForm struct {
	MailType             _const.MailType           `json:"mail_type" example:"plain mail"`
	PostFee              float32                   `json:"post_fee" example:"1.2"`
	PostFeePaymentType   _const.PostFeePaymentType `json:"post_fee_payment_type" example:"stamps"`
	PostRegion           _const.PostRegion         `json:"post_region"`
	RecipientType        _const.RecipientType      `json:"recipient_type" example:"friend"`
	RecipientName        string                    `json:"recipient_name" example:"Jimmy Carter"`
	RecipientZipCode     string                    `json:"recipient_zip_code" example:"610061"`
	RecipientDistrict    string                    `json:"recipient_district"`
	RecipientAddress     string                    `json:"recipient_address"`
	RecipientPhoneNumber string                    `json:"recipient_phone_number"`
}

func (_this *SendMailForm) ToM() *m.SentMail {
	return &m.SentMail{
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
