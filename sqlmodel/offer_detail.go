package sqlmodel

const TableNameOfferDetail = "offer_detail"

var OfferDetailColumns = struct {
	AuditStatus    string
	CheckRound     string
	Commission     string
	CreateAt       string
	ID             string
	OfferNo        string
	OfferTitle     string
	OfferType      string
	PayAt          string
	PayStatus      string
	Price          string
	ProductName    string
	Status         string
	SubmitTimeout  string
	TimesLimitType string
	TotalStock     string
	UID            string
	UpdateAt       string
	VoucherNum     string
}{
	AuditStatus:    "`audit_status`",
	CheckRound:     "`check_round`",
	Commission:     "`commission`",
	CreateAt:       "`create_at`",
	ID:             "`id`",
	OfferNo:        "`offer_no`",
	OfferTitle:     "`offer_title`",
	OfferType:      "`offer_type`",
	PayAt:          "`pay_at`",
	PayStatus:      "`pay_status`",
	Price:          "`price`",
	ProductName:    "`product_name`",
	Status:         "`status`",
	SubmitTimeout:  "`submit_timeout`",
	TimesLimitType: "`times_limit_type`",
	TotalStock:     "`total_stock`",
	UID:            "`uid`",
	UpdateAt:       "`update_at`",
	VoucherNum:     "`voucher_num`",
}

type OfferDetail struct {
	AuditStatus    int32  `json:"audit_status"`     //审核状态 0、审核中 1、审核通过 2、审核失败
	CheckRound     int64  `json:"check_round"`      //审核周期
	Commission     int64  `json:"commission"`       //佣金(分)
	CreateAt       int64  `json:"create_at"`        //创建时间
	ID             int64  `json:"id"`               //
	OfferNo        int64  `json:"offer_no"`         //任务编号
	OfferTitle     string `json:"offer_title"`      //任务名称
	OfferType      int32  `json:"offer_type"`       //任务类型 对应offer_types表
	PayAt          int64  `json:"pay_at"`           //支付时间
	PayStatus      int32  `json:"pay_status"`       //支付状态 0、未支付  1、已支付
	Price          int64  `json:"price"`            //任务单价(分)
	ProductName    string `json:"product_name"`     //产品名称
	Status         int32  `json:"status"`           //任务状态 0、未启动 1、开始中 2、已停用
	SubmitTimeout  int64  `json:"submit_timeout"`   //提交时限
	TimesLimitType int32  `json:"times_limit_type"` //任务次数限制类型  0不限 1仅一次 2每日一次
	TotalStock     int64  `json:"total_stock"`      //任务总量
	UID            int64  `json:"uid"`              //发布用户ID
	UpdateAt       int64  `json:"update_at"`        //更新时间
	VoucherNum     int32  `json:"voucher_num"`      //所需凭证数量
}

// TableName OfferDetail's table name
func (*OfferDetail) TableName() string {
	return TableNameOfferDetail
}
