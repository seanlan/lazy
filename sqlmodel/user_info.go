package sqlmodel

const TableNameUserInfo = "user_info"

var UserInfoColumns = struct {
	Avatar     string
	City       string
	Country    string
	CreateAt   string
	ID         string
	IsFrozen   string
	Nickname   string
	Phone      string
	Province   string
	RegisterIP string
	Sex        string
	UID        string
	UpdateAt   string
}{
	Avatar:     "`avatar`",
	City:       "`city`",
	Country:    "`country`",
	CreateAt:   "`create_at`",
	ID:         "`id`",
	IsFrozen:   "`is_frozen`",
	Nickname:   "`nickname`",
	Phone:      "`phone`",
	Province:   "`province`",
	RegisterIP: "`register_ip`",
	Sex:        "`sex`",
	UID:        "`uid`",
	UpdateAt:   "`update_at`",
}

type UserInfo struct {
	Avatar     string `json:"avatar"`      //用户头像
	City       string `json:"city"`        //城市
	Country    string `json:"country"`     //国家(地区)
	CreateAt   int64  `json:"create_at"`   //创建时间
	ID         int64  `json:"id"`          //
	IsFrozen   int32  `json:"is_frozen"`   //是否冻结
	Nickname   string `json:"nickname"`    //用户昵称
	Phone      string `json:"phone"`       //用户手机号
	Province   string `json:"province"`    //省份
	RegisterIP string `json:"register_ip"` //注册IP
	Sex        int32  `json:"sex"`         //用户性别
	UID        int64  `json:"uid"`         //用户ID
	UpdateAt   int64  `json:"update_at"`   //更新时间
}

// TableName UserInfo's table name
func (*UserInfo) TableName() string {
	return TableNameUserInfo
}
