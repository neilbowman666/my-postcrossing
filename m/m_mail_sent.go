package m

import "gorm.io/gorm"

type SentMail struct {
	gorm.Model `json:"-"`

	MailType MailType

	PostFee            float32
	PostFeePaymentType PostFeePaymentType

	PostRegion PostRegion

	RecipientType        string
	RecipientName        string
	RecipientZipCode     string
	RecipientDistrict    string
	RecipientAddress     string
	RecipientPhoneNumber string
}
