package generator

import (
	"github.com/jimsmart/schema"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
)

type GormModelGenerator struct {
	DB          *gorm.DB
	PackageName string
	Database    string
	OutPath     string
	TmplPaht    string
}

func NewGormGenerator(connStr, database, packageName, tmplPath, outPath string) *GormModelGenerator {
	if !strings.HasSuffix(tmplPath, "/") {
		tmplPath += "/"
	}
	if !strings.HasSuffix(outPath, "/") {
		outPath += "/"
	}
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &GormModelGenerator{
		DB:          db,
		PackageName: packageName,
		Database:    database,
		OutPath:     outPath,
		TmplPaht:    tmplPath,
	}
}

// Tables 获取所有数据表 不包含分表
func (g *GormModelGenerator) Tables() (dbTables []string) {
	sqlDB, _ := g.DB.DB()
	schemaTables, _ := schema.TableNames(sqlDB)
	for _, st := range schemaTables {
		isShard := false
		s := strings.Split(st[1], "_")
		if len(s) > 0 {
			_, err := strconv.ParseInt(s[len(s)-1], 10, 64)
			if err == nil {
				isShard = true
			}
		}
		if st[0] == g.Database && !isShard {
			dbTables = append(dbTables, st[1])
		}
	}
	return
}

func (g *GormModelGenerator) GenTableStruct(tableName string) BaseStruct {
	base := BaseStruct{
		Package:    g.PackageName,
		StructName: g.DB.NamingStrategy.SchemaName(tableName),
		TableName:  tableName,
	}
	cols, _ := getTbColumns(g.DB, g.Database, tableName)
	for _, col := range cols {
		zap.S().Info(col.ColumnType)
		m := toMember(col)
		m.Name = g.DB.NamingStrategy.SchemaName(m.Name)
		base.Members = append(base.Members, m)

	}
	return base
}

func (g *GormModelGenerator) Models() (models []BaseStruct) {
	for _, table := range g.Tables() {
		models = append(models, g.GenTableStruct(table))
	}
	return
}

func (g *GormModelGenerator) Gen() {
	var models []BaseStruct
	for _, table := range g.Tables() {
		models = append(models, g.GenTableStruct(table))
	}
	os.MkdirAll(g.OutPath, os.ModePerm)
	for _, model := range g.Models() {
		outFile := g.OutPath + model.TableName + ".go"
		tmplFile := g.TmplPaht + "gorm_model.tmpl"
		render(outFile, tmplFile, model, true)
	}
}
