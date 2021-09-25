package sqlmodel

const TableNameOfferPayment = "offer_payment"

var OfferPaymentColumns = struct {
	Amount     string
	Balance    string
	Discount   string
	ID         string
	OfferNo    string
	OutTradeNo string
	PayAmount  string
	PayAt      string
	PayChannel string
	PayStatus  string
	PayTradeNo string
}{
	Amount:     "`amount`",
	Balance:    "`balance`",
	Discount:   "`discount`",
	ID:         "`id`",
	OfferNo:    "`offer_no`",
	OutTradeNo: "`out_trade_no`",
	PayAmount:  "`pay_amount`",
	PayAt:      "`pay_at`",
	PayChannel: "`pay_channel`",
	PayStatus:  "`pay_status`",
	PayTradeNo: "`pay_trade_no`",
}

type OfferPayment struct {
	Amount     int64  `json:"amount"`       //所需要的总金额（分）
	Balance    int64  `json:"balance"`      //账号余额抵扣（分）
	Discount   int64  `json:"discount"`     //折扣部分（分）
	ID         int64  `json:"id"`           //
	OfferNo    int64  `json:"offer_no"`     //任务编号
	OutTradeNo string `json:"out_trade_no"` //商户支付编号
	PayAmount  int64  `json:"pay_amount"`   //实际支付金额（分）
	PayAt      int64  `json:"pay_at"`       //支付时间
	PayChannel int32  `json:"pay_channel"`  //支付渠道 1微信 2支付宝
	PayStatus  int32  `json:"pay_status"`   //支付状态 0未支付 1已支付
	PayTradeNo string `json:"pay_trade_no"` //支付流水号
}

// TableName OfferPayment's table name
func (*OfferPayment) TableName() string {
	return TableNameOfferPayment
}
