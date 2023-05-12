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
