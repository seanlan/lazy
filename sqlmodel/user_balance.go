package sqlmodel

const TableNameUserBalance = "user_balance"

var UserBalanceColumns = struct {
	BalanceAsset  string
	BalanceIncome string
	CreateAt      string
	ID            string
	Income        string
	IncomeToday   string
	InviteCount   string
	Level         string
	Recharge      string
	Score         string
	TaskFinished  string
	UID           string
}{
	BalanceAsset:  "`balance_asset`",
	BalanceIncome: "`balance_income`",
	CreateAt:      "`create_at`",
	ID:            "`id`",
	Income:        "`income`",
	IncomeToday:   "`income_today`",
	InviteCount:   "`invite_count`",
	Level:         "`level`",
	Recharge:      "`recharge`",
	Score:         "`score`",
	TaskFinished:  "`task_finished`",
	UID:           "`uid`",
}

type UserBalance struct {
	BalanceAsset  int64 `json:"balance_asset"`  //用户资金账户余额（分）
	BalanceIncome int64 `json:"balance_income"` //用户收入账户余额（分）
	CreateAt      int64 `json:"create_at"`      //创建时间
	ID            int64 `json:"id"`             //
	Income        int64 `json:"income"`         //用户累计收入（分）
	IncomeToday   int64 `json:"income_today"`   //用户今日收入（分）
	InviteCount   int64 `json:"invite_count"`   //用户邀请人数
	Level         int32 `json:"level"`          //用户等级
	Recharge      int64 `json:"recharge"`       //用户累计充值额（分）
	Score         int32 `json:"score"`          //用户积分
	TaskFinished  int64 `json:"task_finished"`  //用户完成任务数量
	UID           int64 `json:"uid"`            //用户ID
}

// TableName UserBalance's table name
func (*UserBalance) TableName() string {
	return TableNameUserBalance
}
