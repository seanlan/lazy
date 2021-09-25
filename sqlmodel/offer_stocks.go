package sqlmodel

const TableNameOfferStock = "offer_stocks"

var OfferStockColumns = struct {
	AppealStock  string
	AuditStock   string
	FailedStock  string
	ID           string
	OfferNo      string
	ProcessStock string
	SuccessStock string
	SurplusStock string
	TotalStock   string
	UseStock     string
}{
	AppealStock:  "`appeal_stock`",
	AuditStock:   "`audit_stock`",
	FailedStock:  "`failed_stock`",
	ID:           "`id`",
	OfferNo:      "`offer_no`",
	ProcessStock: "`process_stock`",
	SuccessStock: "`success_stock`",
	SurplusStock: "`surplus_stock`",
	TotalStock:   "`total_stock`",
	UseStock:     "`use_stock`",
}

type OfferStock struct {
	AppealStock  int64 `json:"appeal_stock"`  //申诉中数量
	AuditStock   int64 `json:"audit_stock"`   //待审核数量
	FailedStock  int64 `json:"failed_stock"`  //失败的数量
	ID           int64 `json:"id"`            //
	OfferNo      int64 `json:"offer_no"`      //任务编号
	ProcessStock int64 `json:"process_stock"` //进行中数量
	SuccessStock int64 `json:"success_stock"` //成功的数量
	SurplusStock int64 `json:"surplus_stock"` //剩余量
	TotalStock   int64 `json:"total_stock"`   //总库存
	UseStock     int64 `json:"use_stock"`     //已接取量
}

// TableName OfferStock's table name
func (*OfferStock) TableName() string {
	return TableNameOfferStock
}
