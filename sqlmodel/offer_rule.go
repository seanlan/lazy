package sqlmodel

const TableNameOfferRule = "offer_rule"

var OfferRuleColumns = struct {
	Command   string
	Describe  string
	ID        string
	OfferNo   string
	OfferText string
	OfferURL  string
	Qrcode    string
	RuleType  string
}{
	Command:   "`command`",
	Describe:  "`describe`",
	ID:        "`id`",
	OfferNo:   "`offer_no`",
	OfferText: "`offer_text`",
	OfferURL:  "`offer_url`",
	Qrcode:    "`qrcode`",
	RuleType:  "`rule_type`",
}

type OfferRule struct {
	Command   string `json:"command"`    //口令
	Describe  string `json:"describe"`   //描述
	ID        int64  `json:"id"`         //
	OfferNo   int64  `json:"offer_no"`   //任务编号
	OfferText string `json:"offer_text"` //纯文本
	OfferURL  string `json:"offer_url"`  //网址
	Qrcode    string `json:"qrcode"`     //二维码图片
	RuleType  int32  `json:"rule_type"`  //打开方式 1、二维码 2、网址 3、口令 4、纯文本
}

// TableName OfferRule's table name
func (*OfferRule) TableName() string {
	return TableNameOfferRule
}
