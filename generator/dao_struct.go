package generator

type DaoBaseStruct struct {
	DaoPackageName string
}

type DaoStruct struct {
	StructName       string
	DaoPackageName   string
	ModelPackageName string
	ModelFQPN        string
}
