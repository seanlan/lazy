package sqlmodel

const TableNameUserRelation = "user_relation"

var UserRelationColumns = struct {
	CreateAt string
	FansUID  string
	ID       string
	Leve     string
	Status   string
	UID      string
}{
	CreateAt: "`create_at`",
	FansUID:  "`fans_uid`",
	ID:       "`id`",
	Leve:     "`leve`",
	Status:   "`status`",
	UID:      "`uid`",
}

type UserRelation struct {
	CreateAt int64 `json:"create_at"` //创建时间
	FansUID  int64 `json:"fans_uid"`  //徒弟用户ID
	ID       int64 `json:"id"`        //
	Leve     int32 `json:"leve"`      //关系层级
	Status   int32 `json:"status"`    //有效状态
	UID      int64 `json:"uid"`       //用户ID
}

// TableName UserRelation's table name
func (*UserRelation) TableName() string {
	return TableNameUserRelation
}
