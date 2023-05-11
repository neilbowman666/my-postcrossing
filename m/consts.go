package m

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
	MailTypeLittlePackage  MailType = "post little package"  // 包裹：小包
)

const PostFeePaymentTypeCash = "cash"     // 现金
const PostFeePaymentTypeStamps = "stamps" // 邮票

const PostRegionDomesticLocal = "domestic local"                      // 国内本埠
const PostRegionDomestic = "domestic"                                 // 国内外埠
const PostRegionOverseasByAir = "overseas (by air)"                   // 跨洋航空
const PostRegionOverseasSurface = "overseas (surface)"                // 跨洋水陆路
const PostRegionOverseasByAirAndSurface = "overseas (by air surface)" // 跨洋航空水陆路
