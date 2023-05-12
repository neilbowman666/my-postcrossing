package form

import (
	_const "my-postcrossing/const"
	"my-postcrossing/m"
	"time"
)

type SendMailForm struct {
	SendDate             time.Time                 `json:"send_date" example:"2023-05-10T12:00:00Z"`
	MailType             _const.MailType           `json:"mail_type" example:"plain mail"`
	PostFee              float32                   `json:"post_fee" example:"1.2"`
	PostFeePaymentType   _const.PostFeePaymentType `json:"post_fee_payment_type" example:"stamps"`
	PostRegion           _const.PostRegion         `json:"post_region" example:"domestic local"`
	RecipientType        _const.RecipientType      `json:"recipient_type" example:"friend"`
	RecipientName        string                    `json:"recipient_name" example:"Jimmy Carter"`
	RecipientZipCode     string                    `json:"recipient_zip_code" example:"610061"`
	RecipientDistrict    string                    `json:"recipient_district" example:"Sichaunsheng Chengdushi Qingyangqu"`
	RecipientAddress     string                    `json:"recipient_address" example:"Caotanglu 1 Hao"`
	RecipientPhoneNumber string                    `json:"recipient_phone_number" example:"13900000001"`
}

func (_this *SendMailForm) ToM() *m.SentMail {
	return &m.SentMail{
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
