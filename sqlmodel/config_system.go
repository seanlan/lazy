package sqlmodel

const TableNameConfigSystem = "config_system"

var ConfigSystemColumns = struct {
	ID     string
	Key    string
	Remark string
	Value  string
}{
	ID:     "`id`",
	Key:    "`key`",
	Remark: "`remarks`",
	Value:  "`value`",
}

type ConfigSystem struct {
	ID     int32  `json:"id"`      //
	Key    string `json:"key"`     //配置项
	Remark string `json:"remarks"` //配置描述
	Value  string `json:"value"`   //配置值
}

// TableName ConfigSystem's table name
func (*ConfigSystem) TableName() string {
	return TableNameConfigSystem
}
