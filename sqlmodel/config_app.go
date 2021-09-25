package sqlmodel

const TableNameConfigApp = "config_app"

var ConfigAppColumns = struct {
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

type ConfigApp struct {
	ID     int32  `json:"id"`      //
	Key    string `json:"key"`     //配置项
	Remark string `json:"remarks"` //配置描述
	Value  string `json:"value"`   //配置值
}

// TableName ConfigApp's table name
func (*ConfigApp) TableName() string {
	return TableNameConfigApp
}
