package m

import (
	"gorm.io/gorm"
	"time"
)

type PostmarkCollection struct {
	gorm.Model            `json:"-"`
	PostmarkDescription   string
	PostmarkCountScenic   int
	PostmarkCountDaily    int
	PostmarkScore         int
	PostOfficeName        string
	PostOfficePhoneNumber string
	PostOfficeZipCode     string
	PostOfficeDistrict    string
	PostOfficeAddress     string
	WhenConfirmed         *time.Time
	WhenSent              *time.Time
	WhenReclaimed         *time.Time
}
