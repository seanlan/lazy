package sqlmodel

const TableNameConfigBanner = "config_banner"

var ConfigBannerColumns = struct {
	ID       string
	ImageURL string
	Link     string
	LinkType string
	ViewType string
	Weight   string
}{
	ID:       "`id`",
	ImageURL: "`image_url`",
	Link:     "`link`",
	LinkType: "`link_type`",
	ViewType: "`view_type`",
	Weight:   "`weight`",
}

type ConfigBanner struct {
	ID       int32  `json:"id"`        //
	ImageURL string `json:"image_url"` //图片地址
	Link     string `json:"link"`      //跳转地址
	LinkType int32  `json:"link_type"` //跳转方式
	ViewType int32  `json:"view_type"` //展示位置
	Weight   int32  `json:"weight"`    //排序值
}

// TableName ConfigBanner's table name
func (*ConfigBanner) TableName() string {
	return TableNameConfigBanner
}
