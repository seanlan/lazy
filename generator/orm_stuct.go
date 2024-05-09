package generator

import (
	"bytes"
	"fmt"
	"gorm.io/gorm"
)

const (
	ModelPkg = "model"

	//query table structure
	columnQuery = "SELECT * FROM information_schema.columns WHERE table_schema = ? AND table_name =?"
)

type dataTypeMapping func(detailType string) (finalType string)
type dataTypeMap map[string]dataTypeMapping

var (
	defaultDataType             = "string"
	dataType        dataTypeMap = map[string]dataTypeMapping{
		"numeric":          func(string) string { return "float64" },
		"integer":          func(string) string { return "int32" },
		"int":              func(string) string { return "int32" },
		"int2":             func(string) string { return "int32" },
		"int4":             func(string) string { return "int32" },
		"int8":             func(string) string { return "int64" },
		"smallint":         func(string) string { return "int32" },
		"mediumint":        func(string) string { return "int32" },
		"bigint":           func(string) string { return "int64" },
		"smallserial":      func(string) string { return "int32" },
		"serial":           func(string) string { return "int32" },
		"bigserial":        func(string) string { return "int64" },
		"float":            func(string) string { return "float32" },
		"real":             func(string) string { return "float64" },
		"double":           func(string) string { return "float64" },
		"double precision": func(string) string { return "float64" },
		"decimal":          func(string) string { return "float64" },
		"money":            func(string) string { return "float64" },
		"char":             func(string) string { return "string" },
		"varchar":          func(string) string { return "string" },
		"tinytext":         func(string) string { return "string" },
		"mediumtext":       func(string) string { return "string" },
		"longtext":         func(string) string { return "string" },
		"inet":             func(string) string { return "string" },
		"binary":           func(string) string { return "[]byte" },
		"varbinary":        func(string) string { return "[]byte" },
		"tinyblob":         func(string) string { return "[]byte" },
		"blob":             func(string) string { return "[]byte" },
		"mediumblob":       func(string) string { return "[]byte" },
		"longblob":         func(string) string { return "[]byte" },
		"text":             func(string) string { return "string" },
		"json":             func(string) string { return "string" },
		"enum":             func(string) string { return "string" },
		"time":             func(string) string { return "time.Time" },
		"date":             func(string) string { return "time.Time" },
		"datetime":         func(string) string { return "time.Time" },
		"timestamp":        func(string) string { return "time.Time" },
		"timestamptz":      func(string) string { return "time.Time" },
		"interval":         func(string) string { return "time.Duration" },
		"year":             func(string) string { return "int32" },
		"bit":              func(string) string { return "[]uint8" },
		"boolean":          func(string) string { return "bool" },
		"tinyint":          func(string) string { return "int32" },
		"uuid":             func(string) string { return "uuid.UUID" },
		"text[]":           func(string) string { return "pq.StringArray" },
		"smallint[]":       func(string) string { return "pq.Int32Array" },
		"integer[]":        func(string) string { return "pq.Int32Array" },
		"bigint[]":         func(string) string { return "pq.Int64Array" },
	}
)

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

func buildGormTag(col gorm.ColumnType) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("column:%s;type:%s", col.Name(), col.DatabaseTypeName()))
	if isPrimaryKey, _ := col.PrimaryKey(); isPrimaryKey {
		buf.WriteString(";primaryKey")
		if isAutoIncrement, _ := col.AutoIncrement(); isAutoIncrement {
			buf.WriteString(";autoIncrement")
		}
	} else if nullable, _ := col.Nullable(); !nullable {
		buf.WriteString(";not null")
	}
	if defaultValue, _ := col.DefaultValue(); defaultValue != "" {
		buf.WriteString(fmt.Sprintf(";default:%s", defaultValue))
	}
	return buf.String()
}

func toMemberWithColumnType(col gorm.ColumnType) *Member {
	memberType := dataType.Get(col.DatabaseTypeName(), col.DatabaseTypeName())
	comment, _ := col.Comment()
	defaultValue, _ := col.DefaultValue()
	return &Member{
		Name:          col.Name(),
		Type:          memberType,
		ModelType:     memberType,
		ColumnName:    col.Name(),
		ColumnComment: comment,
		ColumnDefault: defaultValue,
		JSONTag:       col.Name(),
		GORMTag:       buildGormTag(col),
	}
}
