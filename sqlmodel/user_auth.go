package sqlmodel

const TableNameUserAuth = "user_auth"

var UserAuthColumns = struct {
	CreateAt     string
	Credential   string
	ID           string
	Identifier   string
	IdentityType string
	UID          string
	UnionID      string
}{
	CreateAt:     "`create_at`",
	Credential:   "`credential`",
	ID:           "`id`",
	Identifier:   "`identifier`",
	IdentityType: "`identity_type`",
	UID:          "`uid`",
	UnionID:      "`union_id`",
}

type UserAuth struct {
	CreateAt     int64  `json:"create_at"`     //创建时间
	Credential   string `json:"credential"`    //识别密钥
	ID           int64  `json:"id"`            //
	Identifier   string `json:"identifier"`    //识别码
	IdentityType int32  `json:"identity_type"` //识别类型 1微信app 2设备号 3手机号
	UID          int64  `json:"uid"`           //用户ID
	UnionID      string `json:"union_id"`      //用户union_id
}

// TableName UserAuth's table name
func (*UserAuth) TableName() string {
	return TableNameUserAuth
}
