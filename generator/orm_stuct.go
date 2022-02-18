package generator

import (
	"bytes"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

const (
	ModelPkg = "model"

	//query table structure
	columnQuery = "SELECT * FROM information_schema.columns WHERE table_schema = ? AND table_name =?"
)

var (
	defaultDataType             = "string"
	dataType        dataTypeMap = map[string]func(detailType string) string{
		"int": func(ct string) string {
			if strings.Contains(ct, "unsigned") {
				return "uint32"
			} else {
				return "int32"
			}
		},
		"integer": func(ct string) string {
			if strings.Contains(ct, "unsigned") {
				return "uint32"
			} else {
				return "int32"
			}
		},
		"smallint": func(ct string) string {
			if strings.Contains(ct, "unsigned") {
				return "uint32"
			} else {
				return "int32"
			}
		},
		"mediumint": func(ct string) string {
			if strings.Contains(ct, "unsigned") {
				return "uint32"
			} else {
				return "int32"
			}
		},
		"bigint": func(ct string) string {
			if strings.Contains(ct, "unsigned") {
				return "uint64"
			} else {
				return "int64"
			}
		},
		"float":      func(ct string) string { return "float32" },
		"double":     func(ct string) string { return "float64" },
		"decimal":    func(ct string) string { return "float64" },
		"char":       func(ct string) string { return "string" },
		"varchar":    func(ct string) string { return "string" },
		"tinytext":   func(ct string) string { return "string" },
		"mediumtext": func(ct string) string { return "string" },
		"longtext":   func(ct string) string { return "string" },
		"tinyblob":   func(ct string) string { return "[]byte" },
		"blob":       func(ct string) string { return "[]byte" },
		"mediumblob": func(ct string) string { return "[]byte" },
		"longblob":   func(ct string) string { return "[]byte" },
		"text":       func(ct string) string { return "string" },
		"json":       func(ct string) string { return "string" },
		"enum":       func(ct string) string { return "string" },
		"time":       func(ct string) string { return "time.Time" },
		"date":       func(ct string) string { return "time.Time" },
		"datetime":   func(ct string) string { return "time.Time" },
		"timestamp":  func(ct string) string { return "time.Time" },
		"year":       func(ct string) string { return "int32" },
		"bit":        func(ct string) string { return "[]uint8" },
		"boolean":    func(ct string) string { return "bool" },
		"tinyint": func(detailType string) string {
			//if strings.HasPrefix(detailType, "tinyint(1)") {
			//	return "bool"
			//}
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
	ColumnDefault string `gorm:"column:COLUMN_DEFAULT"`
	Extra         string `gorm:"column:EXTRA"`
	IsNullable    string `gorm:"column:IS_NULLABLE"`
}

func (c *Column) IsPrimaryKey() bool {
	return c != nil && c.ColumnKey == "PRI"
}

func (c *Column) AutoIncrement() bool {
	return c != nil && c.Extra == "auto_increment"
}

func (c *Column) buildGormTag() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("column:%s;type:%s", c.ColumnName, c.ColumnType))
	if c.IsPrimaryKey() {
		buf.WriteString(";primaryKey")
		buf.WriteString(fmt.Sprintf(";autoIncrement:%t", c.AutoIncrement()))
	} else if c.IsNullable != "YES" {
		buf.WriteString(";not null")
	}
	if c.ColumnDefault != "" {
		buf.WriteString(fmt.Sprintf(";default:%s", c.ColumnDefault))
	}
	return buf.String()
}

// Member user input structures
type Member struct {
	Name          string
	Type          string
	NewType       string
	ColumnName    string
	ColumnComment string
	ColumnDefault string
	ModelType     string
	JSONTag       string
	GORMTag       string
	NewTag        string
}

type BaseStruct struct {
	Package    string
	StructName string
	TableName  string
	Imports    string
	Members    []*Member
}

func getTbColumns(db *gorm.DB, schemaName string, tableName string) (result []*Column, err error) {
	err = db.Raw(columnQuery, schemaName, tableName).Scan(&result).Error
	return
}

func toMember(field *Column) *Member {
	memberType := dataType.Get(field.DataType, field.ColumnType)
	return &Member{
		Name:          field.ColumnName,
		Type:          memberType,
		ModelType:     memberType,
		ColumnName:    field.ColumnName,
		ColumnComment: field.ColumnComment,
		ColumnDefault: field.ColumnDefault,
		JSONTag:       field.ColumnName,
		GORMTag:       field.buildGormTag(),
	}
}
