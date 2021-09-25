package sqlmodel

const TableNameOfferRecord = "offer_records"

var OfferRecordColumns = struct {
	AppealStatus string
	AppealTime   string
	Commission   string
	CreateAt     string
	EmployerUID  string
	EndAt        string
	EndStatus    string
	ID           string
	OfferNo      string
	OfferTitle   string
	OfferType    string
	ProductName  string
	RecordNo     string
	Status       string
	SubmitAt     string
	SubmitTime   string
	TimeoutAt    string
	UID          string
	UpdateAt     string
}{
	AppealStatus: "`appeal_status`",
	AppealTime:   "`appeal_times`",
	Commission:   "`commission`",
	CreateAt:     "`create_at`",
	EmployerUID:  "`employer_uid`",
	EndAt:        "`end_at`",
	EndStatus:    "`end_status`",
	ID:           "`id`",
	OfferNo:      "`offer_no`",
	OfferTitle:   "`offer_title`",
	OfferType:    "`offer_type`",
	ProductName:  "`product_name`",
	RecordNo:     "`record_no`",
	Status:       "`status`",
	SubmitAt:     "`submit_at`",
	SubmitTime:   "`submit_times`",
	TimeoutAt:    "`timeout_at`",
	UID:          "`uid`",
	UpdateAt:     "`update_at`",
}

type OfferRecord struct {
	AppealStatus int32  `json:"appeal_status"` //申述状态 0未提交 1申诉中 2申诉成功 3申诉失败
	AppealTime   int32  `json:"appeal_times"`  //申诉次数
	Commission   int64  `json:"commission"`    //获得的佣金
	CreateAt     int64  `json:"create_at"`     //领取任务时间
	EmployerUID  int64  `json:"employer_uid"`  //雇佣者用户编号
	EndAt        int64  `json:"end_at"`        //完结到期时间
	EndStatus    int32  `json:"end_status"`    //完结状态 0 未完结 1已完结
	ID           int64  `json:"id"`            //
	OfferNo      int64  `json:"offer_no"`      //任务编号
	OfferTitle   string `json:"offer_title"`   //任务标题
	OfferType    int32  `json:"offer_type"`    //任务类型
	ProductName  string `json:"product_name"`  //产品名称
	RecordNo     int64  `json:"record_no"`     //记录编号
	Status       int32  `json:"status"`        //状态 0未提交 1审核中 2完成 3失败
	SubmitAt     int64  `json:"submit_at"`     //提交时间
	SubmitTime   int32  `json:"submit_times"`  //提交次数（第几次提交）
	TimeoutAt    int64  `json:"timeout_at"`    //提交超时时间
	UID          int64  `json:"uid"`           //用户编号
	UpdateAt     int64  `json:"update_at"`     //更新时间
}

// TableName OfferRecord's table name
func (*OfferRecord) TableName() string {
	return TableNameOfferRecord
}
