package form

import (
	"my-postcrossing/m"
	"time"
)

type AddPostmarkCollectionForm struct {
	PostmarkDescription   string `json:"postmark_description" example:"2枚风景戳+1枚日戳"`
	PostmarkCountScenic   int    `json:"postmark_count_scenic" example:"2"`
	PostmarkCountDaily    int    `json:"postmark_count_daily" example:"1"`
	PostmarkScore         int    `json:"postmark_score" example:"76"`
	PostOfficeName        string `json:"post_office_name" example:"天上西藏主题邮局"`
	PostOfficePhoneNumber string `json:"post_office_phone_number" example:"0891-75660233"`
	PostOfficeZipCode     string `json:"post_office_zip_code" example:"850030"`
	PostOfficeDistrict    string `json:"post_office_district" example:"西藏自治区 拉萨市 城关区"`
	PostOfficeAddress     string `json:"post_office_address" example:"北京中路33号"`
	WhenConfirmed         *int64 `json:"when_confirmed" example:"1684310643"`
	WhenSent              *int64 `json:"when_sent" example:"1684310643"`
	WhenReclaimed         *int64 `json:"when_reclaimed" example:"1685355405"`
}

func (f AddPostmarkCollectionForm) ToM() (*m.PostmarkCollection, error) {
	var whenConfirmed *time.Time = nil
	if f.WhenConfirmed != nil {
		_whenConfirmed := time.Unix(*f.WhenConfirmed, 0)
		whenConfirmed = &_whenConfirmed
	}
	var whenSent *time.Time = nil
	if f.WhenSent != nil {
		_whenSent := time.Unix(*f.WhenSent, 0)
		whenSent = &_whenSent
	}
	var whenReclaimed *time.Time = nil
	if f.WhenReclaimed != nil {
		_whenReclaimed := time.Unix(*f.WhenReclaimed, 0)
		whenReclaimed = &_whenReclaimed
	}

	return &m.PostmarkCollection{
		PostmarkDescription:   f.PostmarkDescription,
		PostmarkCountScenic:   f.PostmarkCountScenic,
		PostmarkCountDaily:    f.PostmarkCountDaily,
		PostmarkScore:         f.PostmarkScore,
		PostOfficeName:        f.PostOfficeName,
		PostOfficePhoneNumber: f.PostOfficePhoneNumber,
		PostOfficeZipCode:     f.PostOfficeZipCode,
		PostOfficeDistrict:    f.PostOfficeDistrict,
		PostOfficeAddress:     f.PostOfficeAddress,
		WhenConfirmed:         whenConfirmed,
		WhenSent:              whenSent,
		WhenReclaimed:         whenReclaimed,
	}, nil
}
