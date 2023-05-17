package _const

type MailType string

const (
	MailTypePlainMail      MailType = "plain mail"      // 信件：一般
	MailTypeRegisteredMail MailType = "registered mail" // 信件：挂号

	MailTypePlainPostcard      MailType = "plain postcard"      // 明信片：一般
	MailTypeRegisteredPostcard MailType = "registered postcard" // 明信片：挂号

	MailTypePlainPresswork      MailType = "plain presswork"      // 印刷品：一般
	MailTypeRegisteredPresswork MailType = "registered presswork" // 印刷品：挂号

	MailTypePackage        MailType = "post package"         // 包裹：一般
	MailTypeExpressPackage MailType = "post express package" // 包裹：快递
	MailTypeSmallPackage   MailType = "post small package"   // 包裹：小包
)

type PostFeePaymentType string

const (
	PostFeePaymentTypeCash   PostFeePaymentType = "cash"   // 现金
	PostFeePaymentTypeStamps PostFeePaymentType = "stamps" // 邮票
)

type PostRegion string

const (
	PostRegionDomesticLocal           PostRegion = "domestic local"            // 国内本埠
	PostRegionDomestic                PostRegion = "domestic"                  // 国内外埠
	PostRegionOverseasByAir           PostRegion = "overseas (by air)"         // 跨洋航空
	PostRegionOverseasSurface         PostRegion = "overseas (surface)"        // 跨洋水陆路
	PostRegionOverseasByAirAndSurface PostRegion = "overseas (by air surface)" // 跨洋航空水陆路
)

type RecipientType string

const (
	RecipientTypeMyself             RecipientType = "myself"
	RecipientTypeCouple             RecipientType = "couple"
	RecipientTypeRelative           RecipientType = "relative"
	RecipientTypeFriend             RecipientType = "friend"
	RecipientTypePenPal             RecipientType = "pen pal"
	RecipientTypePostmarkCollecting RecipientType = "postmark collecting"
	RecipientTypeOther              RecipientType = "other"
)
