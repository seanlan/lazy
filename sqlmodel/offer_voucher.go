package sqlmodel

const TableNameOfferVoucher = "offer_voucher"

var OfferVoucherColumns = struct {
	Describe string
	ID       string
	Image    string
	Index    string
	OfferNo  string
}{
	Describe: "`describe`",
	ID:       "`id`",
	Image:    "`image`",
	Index:    "`index`",
	OfferNo:  "`offer_no`",
}

type OfferVoucher struct {
	Describe string `json:"describe"` //凭证描述
	ID       int64  `json:"id"`       //
	Image    string `json:"image"`    //凭证示例图
	Index    int32  `json:"index"`    //步骤顺序编号
	OfferNo  int64  `json:"offer_no"` //任务编号
}

// TableName OfferVoucher's table name
func (*OfferVoucher) TableName() string {
	return TableNameOfferVoucher
}
