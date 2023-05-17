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
	WhenConfirmed         string `json:"when_confirmed" example:"2023-05-17T16:04:03+0800"`
	WhenSent              string `json:"when_sent" example:"2023-05-17T16:04:03+0800"`
	WhenReclaimed         string `json:"when_reclaimed" example:"2023-05-29T18:16:45+0800"`
}

func (f AddPostmarkCollectionForm) ToM() (*m.PostmarkCollection, error) {
	var whenConfirmed *time.Time = nil
	if f.WhenConfirmed != "" {
		_whenConfirmed, err := time.Parse(time.RFC3339, f.WhenConfirmed)
		if err != nil {
			return nil, err
		}
		whenConfirmed = &_whenConfirmed
	}
	var whenSent *time.Time = nil
	if f.WhenSent != "" {
		_whenSent, err := time.Parse(time.RFC3339, f.WhenSent)
		if err != nil {
			return nil, err
		}
		whenSent = &_whenSent
	}
	var whenReclaimed *time.Time = nil
	if f.WhenReclaimed != "" {
		_whenReclaimed, err := time.Parse(time.RFC3339, f.WhenReclaimed)
		if err != nil {
			return nil, err
		}
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
