package sqlmodel

const TableNameOfferRecordsVoucher = "offer_records_voucher"

var OfferRecordsVoucherColumns = struct {
	Describe  string
	ID        string
	ImageHash string
	Image     string
	Index     string
	OfferNo   string
	RecordNo  string
	UID       string
}{
	Describe:  "`describe`",
	ID:        "`id`",
	ImageHash: "`image_hash`",
	Image:     "`images`",
	Index:     "`index`",
	OfferNo:   "`offer_no`",
	RecordNo:  "`record_no`",
	UID:       "`uid`",
}

type OfferRecordsVoucher struct {
	Describe  string `json:"describe"`   //凭证说明
	ID        int64  `json:"id"`         //
	ImageHash string `json:"image_hash"` //凭证的特征值
	Image     string `json:"images"`     //凭证图片
	Index     int32  `json:"index"`      //步骤顺序编号
	OfferNo   int64  `json:"offer_no"`   //任务编号
	RecordNo  int64  `json:"record_no"`  //记录编号
	UID       int64  `json:"uid"`        //用户编号
}

// TableName OfferRecordsVoucher's table name
func (*OfferRecordsVoucher) TableName() string {
	return TableNameOfferRecordsVoucher
}
