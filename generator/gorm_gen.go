package generator

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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

func parseDSN(url string) (string, string, error) {
	a := strings.SplitN(url, "://", 2)
	if len(a) != 2 {
		return "", "", fmt.Errorf(`failed to parse dsn: "%s"`, url)
	}
	return a[0], a[1], nil
}

func NewGormGenerator(connStr, database, prefix, packageName, tmplPath, modelPackage, modelPath, daoPackage, daoPath string) *GormDaoGenerator {
	var (
		dialect gorm.Dialector
	)
	scheme, dail, _err := parseDSN(connStr)
	if _err != nil {
		zap.S().Infof("parse dsn error : %s %v", connStr, _err)
		return nil
	}
	switch scheme {
	case "mysql":
		dialect = mysql.Open(dail)
	case "postgres":
		dialect = postgres.Open(connStr)
	}
	db, err := gorm.Open(dialect, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: prefix, SingularTable: true,
		},
	})
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
	schemaTables, err := g.DB.Migrator().GetTables()
	if err != nil {
		zap.S().Info("get table names error : ", err)
		return
	}
	for _, st := range schemaTables {
		isShard := false
		s := strings.Split(st, "_")
		if len(s) > 0 {
			_, err = strconv.ParseInt(s[len(s)-1], 10, 64)
			if err == nil {
				isShard = true
			}
		}
		if isShard {
			continue
		}
		dbTables = append(dbTables, st)
	}
	return
}

func (g *GormDaoGenerator) GenTableStruct(tableName string) BaseStruct {
	base := BaseStruct{
		Package:    g.ModelPackageName,
		StructName: g.DB.NamingStrategy.SchemaName(tableName),
		TableName:  tableName,
	}
	var imports = make([]string, 0)
	columns, _ := g.DB.Migrator().ColumnTypes(tableName)
	for _, col := range columns {
		m := toMemberWithColumnType(col)
		m.Name = g.DB.NamingStrategy.SchemaName(m.Name)
		base.Members = append(base.Members, m)
		if strings.Contains(m.ModelType, "time.") {
			imports = append(imports, "\"time\"")
		} else if strings.Contains(m.ModelType, "pq.") {
			imports = append(imports, "\"github.com/lib/pq\"")
		} else if strings.Contains(m.ModelType, "uuid.") {
			imports = append(imports, "\"github.com/google/uuid\"")
		}
	}
	//cols, _ := getTbColumns(g.DB, g.Database, tableName)
	//for _, col := range cols {
	//	m := toMember(col)
	//	m.Name = g.DB.NamingStrategy.SchemaName(m.Name)
	//	base.Members = append(base.Members, m)
	//	if strings.Contains(m.ModelType, "time.") {
	//		imports = append(imports, "\"time\"")
	//	} else if strings.Contains(m.ModelType, "pq.") {
	//		imports = append(imports, "\"github.com/lib/pq\"")
	//	} else if strings.Contains(m.ModelType, "uuid.") {
	//		imports = append(imports, "\"github.com/google/uuid\"")
	//	}
	//}
	base.Imports = "import (\n" + strings.Join(imports, "\n") + "\n)"
	return base
}

func (g *GormDaoGenerator) GenDaoStruct(tableName string) DaoStruct {
	dao := DaoStruct{
		TableName:        tableName,
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
	modelBaseFile := filepath.Join(modelOutPath, "base.go")
	RenderWithStruct(modelBaseFile, g.TmplPath, "dao_gorm_model_base.tmpl",
		BaseStruct{Package: g.ModelPackageName}, true)

	for _, model := range sqlModels {
		outFile := filepath.Join(modelOutPath, model.TableName+".go")
		RenderWithStruct(outFile, g.TmplPath, "dao_gorm_model.tmpl", model, true)
		zap.S().Info("generate model " + outFile + " success !")
	}
	daoOutPath := g.DaoPath
	err = os.MkdirAll(daoOutPath, os.ModePerm)
	if err != nil {
		zap.S().Info("mkdir error : " + daoOutPath)
		return
	}
	daoBaseFile := filepath.Join(daoOutPath, "dao_base.go")
	RenderWithStruct(daoBaseFile, g.TmplPath, "dao_gorm_base.tmpl",
		DaoBaseStruct{g.DaoPackageName}, true)
	zap.S().Info("generate dao " + daoBaseFile + " success !")
	for _, dao := range daoModels {
		daoName := dao.TableName
		outFile := filepath.Join(daoOutPath, daoName+".go")
		RenderWithStruct(outFile, g.TmplPath, "dao_gorm.tmpl", dao, true)
		zap.S().Info("generate dao " + outFile + " success !")
	}
}
