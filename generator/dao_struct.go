package generator

type DaoBaseStruct struct {
	DaoPackageName string
}

type DaoStruct struct {
	TableName        string
	StructName       string
	DaoPackageName   string
	ModelPackageName string
	ModelFQPN        string
}
