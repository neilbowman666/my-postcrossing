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

func (_this PostmarkCollection) ToVo() Visible {
	return PostmarkCollectionVo{
		ID:                    _this.ID,
		PostmarkDescription:   _this.PostmarkDescription,
		PostmarkCountScenic:   _this.PostmarkCountScenic,
		PostmarkCountDaily:    _this.PostmarkCountDaily,
		PostmarkScore:         _this.PostmarkScore,
		PostOfficeName:        _this.PostOfficeName,
		PostOfficePhoneNumber: _this.PostOfficePhoneNumber,
		PostOfficeZipCode:     _this.PostOfficeZipCode,
		PostOfficeDistrict:    _this.PostOfficeDistrict,
		PostOfficeAddress:     _this.PostOfficeAddress,
		WhenConfirmed:         _this.WhenConfirmed,
		WhenSent:              _this.WhenSent,
		WhenReclaimed:         _this.WhenReclaimed,
	}
}

type PostmarkCollectionVo struct {
	ID                    uint       `json:"id"`
	PostmarkDescription   string     `json:"postmark_description"`
	PostmarkCountScenic   int        `json:"postmark_count_scenic"`
	PostmarkCountDaily    int        `json:"postmark_count_daily"`
	PostmarkScore         int        `json:"postmark_score"`
	PostOfficeName        string     `json:"post_office_name"`
	PostOfficePhoneNumber string     `json:"post_office_phone_number"`
	PostOfficeZipCode     string     `json:"post_office_zip_code"`
	PostOfficeDistrict    string     `json:"post_office_district"`
	PostOfficeAddress     string     `json:"post_office_address"`
	WhenConfirmed         *time.Time `json:"when_confirmed"`
	WhenSent              *time.Time `json:"when_sent"`
	WhenReclaimed         *time.Time `json:"when_reclaimed"`
}

func (PostmarkCollectionVo) Visible() {
}
