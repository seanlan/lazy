package sqlmodel

const TableNameOfferRecordsExamine = "offer_records_examine"

var OfferRecordsExamineColumns = struct {
	CreateAt string
	ID       string
	Image    string
	RecordNo string
	Remark   string
	Status   string
	Time     string
	UID      string
}{
	CreateAt: "`create_at`",
	ID:       "`id`",
	Image:    "`images`",
	RecordNo: "`record_no`",
	Remark:   "`remarks`",
	Status:   "`status`",
	Time:     "`times`",
	UID:      "`uid`",
}

type OfferRecordsExamine struct {
	CreateAt int64  `json:"create_at"` //审批时间
	ID       int64  `json:"id"`        //
	Image    string `json:"images"`    //审批配图
	RecordNo int64  `json:"record_no"` //记录编号
	Remark   string `json:"remarks"`   //审批备注
	Status   int32  `json:"status"`    //审批状态 0审核中 1审核成功 2审核失败
	Time     int32  `json:"times"`     //审批次数 第几次审批
	UID      int64  `json:"uid"`       //用户编号
}

// TableName OfferRecordsExamine's table name
func (*OfferRecordsExamine) TableName() string {
	return TableNameOfferRecordsExamine
}
