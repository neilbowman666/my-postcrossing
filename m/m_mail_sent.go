package m

import "gorm.io/gorm"

type SentMail struct {
	gorm.Model `json:"-"`

	MailType string

	PostFee            float32
	PostFeePaymentType string

	PostRegion string

	RecipientType        string
	RecipientName        string
	RecipientDistrict    string
	RecipientAddress     string
	RecipientPhoneNumber string
}
