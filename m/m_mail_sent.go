package m

import (
	"gorm.io/gorm"
	"my-postcrossing/const"
)

type SentMail struct {
	gorm.Model `json:"-"`

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
