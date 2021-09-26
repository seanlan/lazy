package generator

import (
	"github.com/jimsmart/schema"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type GormDaoGenerator struct {
	DB               *gorm.DB
	PackageName      string
	ModelPackageName string
	ModelPath        string
	DaoPackageName   string
	DaoPath          string
	Database         string
	TmplPath         string
}

func NewGormGenerator(connStr, database, packageName, tmplPath, modelPackage, modelPath, daoPackage, daoPath string) *GormDaoGenerator {
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		zap.S().Info("mysql connect failed ")
		return nil
	}
	return &GormDaoGenerator{
		DB:               db,
		PackageName:      packageName,
		ModelPackageName: modelPackage,
		ModelPath:        modelPath,
		DaoPackageName:   daoPackage,
		DaoPath:          daoPath,
		Database:         database,
		TmplPath:         tmplPath,
	}
}

// Tables 获取所有数据表 不包含分表
func (g *GormDaoGenerator) Tables() (dbTables []string) {
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

func (g *GormDaoGenerator) GenTableStruct(tableName string) BaseStruct {
	base := BaseStruct{
		Package:    g.ModelPackageName,
		StructName: g.DB.NamingStrategy.SchemaName(tableName),
		TableName:  tableName,
	}
	cols, _ := getTbColumns(g.DB, g.Database, tableName)
	for _, col := range cols {
		m := toMember(col)
		m.Name = g.DB.NamingStrategy.SchemaName(m.Name)
		base.Members = append(base.Members, m)
	}
	return base
}

func (g *GormDaoGenerator) GenDaoStruct(tableName string) DaoStruct {
	dao := DaoStruct{
		StructName:       g.DB.NamingStrategy.SchemaName(tableName),
		DaoPackageName:   g.DaoPackageName,
		ModelPackageName: g.ModelPackageName,
		ModelFQPN:        filepath.Join(g.PackageName, g.ModelPath),
	}
	return dao
}

func (g *GormDaoGenerator) Gen() {
	var sqlModels []BaseStruct
	var daoModels []DaoStruct
	for _, table := range g.Tables() {
		sqlModels = append(sqlModels, g.GenTableStruct(table))
		daoModels = append(daoModels, g.GenDaoStruct(table))
	}
	modelOutPath := g.ModelPath
	err := os.MkdirAll(modelOutPath, os.ModePerm)
	if err != nil {
		zap.S().Info("mkdir error : " + modelOutPath)
		return
	}
	for _, model := range sqlModels {
		outFile := filepath.Join(modelOutPath, model.TableName+".go")
		render(outFile, g.TmplPath, "gorm_model.tmpl", model, true)
		zap.S().Info("generate model " + outFile + " success !")
	}
	daoOutPath := g.DaoPath
	err = os.MkdirAll(daoOutPath, os.ModePerm)
	if err != nil {
		zap.S().Info("mkdir error : " + daoOutPath)
		return
	}
	daoBaseFile := filepath.Join(daoOutPath, "dao_base.go")
	render(daoBaseFile, g.TmplPath, "dao_gorm_base.tmpl",
		DaoBaseStruct{g.DaoPackageName}, true)
	zap.S().Info("generate dao " + daoBaseFile + " success !")
	for _, dao := range daoModels {
		daoName := g.DB.NamingStrategy.TableName(dao.StructName)
		outFile := filepath.Join(daoOutPath, daoName+".go")
		render(outFile, g.TmplPath, "dao_gorm.tmpl", dao, true)
		zap.S().Info("generate dao " + outFile + " success !")
	}
}
