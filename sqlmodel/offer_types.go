package sqlmodel

const TableNameOfferType = "offer_types"

var OfferTypeColumns = struct {
	ID            string
	OfferTypeIcon string
	OfferTypeName string
	Requirement   string
	Weight        string
}{
	ID:            "`id`",
	OfferTypeIcon: "`offer_type_icon`",
	OfferTypeName: "`offer_type_name`",
	Requirement:   "`requirement`",
	Weight:        "`weight`",
}

type OfferType struct {
	ID            int32  `json:"id"`              //
	OfferTypeIcon string `json:"offer_type_icon"` //任务类型图标
	OfferTypeName string `json:"offer_type_name"` //任务类型名称
	Requirement   string `json:"requirement"`     //需要达到的要求
	Weight        int32  `json:"weight"`          //排序权重
}

// TableName OfferType's table name
func (*OfferType) TableName() string {
	return TableNameOfferType
}
