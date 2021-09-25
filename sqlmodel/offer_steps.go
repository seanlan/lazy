package sqlmodel

const TableNameOfferStep = "offer_steps"

var OfferStepColumns = struct {
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

type OfferStep struct {
	Describe string `json:"describe"` //任务步骤描述
	ID       int64  `json:"id"`       //
	Image    string `json:"image"`    //任务步骤图片
	Index    int32  `json:"index"`    //步骤顺序编号
	OfferNo  int64  `json:"offer_no"` //任务编号
}

// TableName OfferStep's table name
func (*OfferStep) TableName() string {
	return TableNameOfferStep
}
