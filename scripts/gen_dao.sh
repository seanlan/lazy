dbuser="root"
dbpwd="q145145145"
dbhost="127.0.0.1:3306"
dbname="fish"
conn="$dbuser:$dbpwd@tcp($dbhost)/$dbname?parseTime=true&loc=Local&charset=utf8mb4&collation=utf8mb4_unicode_ci"
package="fish"
prefix="t_"
modelPackage="sqlmodel"
modelPath="internal/dao/sqlmodel"
daoPackage="dao"
daoPath="internal/dao"
go run main.go dao --conn=$conn --database=$dbname --prefix=$prefix --package=$package --template=$template \
 --model=$modelPackage --model-path=$modelPath --dao=$daoPackage --dao-path=$daoPath
