package sqlmodel

const TableNameOfferRecordsAppeal = "offer_records_appeal"

var OfferRecordsAppealColumns = struct {
	AppealAt string
	CreateAt string
	ID       string
	Image    string
	Reason   string
	RecordNo string
	Remark   string
	Status   string
	Time     string
	UID      string
}{
	AppealAt: "`appeal_at`",
	CreateAt: "`create_at`",
	ID:       "`id`",
	Image:    "`images`",
	Reason:   "`reasons`",
	RecordNo: "`record_no`",
	Remark:   "`remarks`",
	Status:   "`status`",
	Time:     "`times`",
	UID:      "`uid`",
}

type OfferRecordsAppeal struct {
	AppealAt int64  `json:"appeal_at"` //申诉结束时间
	CreateAt int64  `json:"create_at"` //提交申诉时间
	ID       int64  `json:"id"`        //
	Image    string `json:"images"`    //申诉结果图片
	Reason   string `json:"reasons"`   //申诉理由
	RecordNo int64  `json:"record_no"` //记录编号
	Remark   string `json:"remarks"`   //申诉结果备注
	Status   int32  `json:"status"`    //申诉结果 0申诉中 1通过 2失败
	Time     int32  `json:"times"`     //申诉次数
	UID      int64  `json:"uid"`       //用户编号
}

// TableName OfferRecordsAppeal's table name
func (*OfferRecordsAppeal) TableName() string {
	return TableNameOfferRecordsAppeal
}
