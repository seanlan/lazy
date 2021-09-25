package generator

import (
	"gorm.io/gorm"
	"strings"
)

const (
	ModelPkg = "model"

	//query table structure
	columnQuery = "SELECT COLUMN_NAME ,COLUMN_COMMENT ,DATA_TYPE ,IS_NULLABLE ,COLUMN_KEY,COLUMN_TYPE,EXTRA" +
		" FROM information_schema.columns WHERE table_schema = ? AND table_name =?"
)

var (
	defaultDataType             = "string"
	dataType        dataTypeMap = map[string]func(detailType string) string{
		"int":        func(string) string { return "int32" },
		"integer":    func(string) string { return "int32" },
		"smallint":   func(string) string { return "int32" },
		"mediumint":  func(string) string { return "int32" },
		"bigint":     func(string) string { return "int64" },
		"float":      func(string) string { return "float32" },
		"double":     func(string) string { return "float64" },
		"decimal":    func(string) string { return "float64" },
		"char":       func(string) string { return "string" },
		"varchar":    func(string) string { return "string" },
		"tinytext":   func(string) string { return "string" },
		"mediumtext": func(string) string { return "string" },
		"longtext":   func(string) string { return "string" },
		"tinyblob":   func(string) string { return "[]byte" },
		"blob":       func(string) string { return "[]byte" },
		"mediumblob": func(string) string { return "[]byte" },
		"longblob":   func(string) string { return "[]byte" },
		"text":       func(string) string { return "string" },
		"json":       func(string) string { return "string" },
		"enum":       func(string) string { return "string" },
		"time":       func(string) string { return "time.Time" },
		"date":       func(string) string { return "time.Time" },
		"datetime":   func(string) string { return "time.Time" },
		"timestamp":  func(string) string { return "time.Time" },
		"year":       func(string) string { return "int32" },
		"bit":        func(string) string { return "[]uint8" },
		"boolean":    func(string) string { return "bool" },
		"tinyint": func(detailType string) string {
			if strings.HasPrefix(detailType, "tinyint(1)") {
				return "bool"
			}
			return "int32"
		},
	}
)

type dataTypeMap map[string]func(string) string

func (m dataTypeMap) Get(dataType, detailType string) string {
	if convert, ok := m[dataType]; ok {
		return convert(detailType)
	}
	return defaultDataType
}

// Column table column's info
type Column struct {
	TableName     string `gorm:"column:TABLE_NAME"`
	ColumnName    string `gorm:"column:COLUMN_NAME"`
	ColumnComment string `gorm:"column:COLUMN_COMMENT"`
	DataType      string `gorm:"column:DATA_TYPE"`
	ColumnKey     string `gorm:"column:COLUMN_KEY"`
	ColumnType    string `gorm:"column:COLUMN_TYPE"`
	Extra         string `gorm:"column:EXTRA"`
}

// Member user input structures
type Member struct {
	Name          string
	Type          string
	NewType       string
	ColumnName    string
	ColumnComment string
	ModelType     string
	JSONTag       string
	NewTag        string
}

type BaseStruct struct {
	Package    string
	StructName string
	TableName  string
	Members    []*Member
}

func getTbColumns(db *gorm.DB, schemaName string, tableName string) (result []*Column, err error) {
	err = db.Raw(columnQuery, schemaName, tableName).Scan(&result).Error
	return
}

func toMember(field *Column) *Member {
	memberType := dataType.Get(field.DataType, field.ColumnType)
	if field.ColumnName == "deleted_at" && memberType == "time.Time" {
		memberType = "gorm.DeletedAt"
	}
	return &Member{
		Name:          field.ColumnName,
		Type:          memberType,
		ModelType:     memberType,
		ColumnName:    field.ColumnName,
		ColumnComment: field.ColumnComment,
		JSONTag:       field.ColumnName,
	}
}
