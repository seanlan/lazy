dbuser="root"
dbpwd="q145145145"
dbhost="127.0.0.1:3306"
dbname="lucky"
conn="$dbuser:$dbpwd@tcp($dbhost)/$dbname?parseTime=true&loc=Local&charset=utf8mb4&collation=utf8mb4_unicode_ci"
package="github.com/seanlan/lazy"
template=""
modelPackage="sqlmodel"
modelPath="app/dao/sqlmodel"
daoPackage="dao"
daoPath="app/dao"
go run main.go dao --conn=$conn --database=$dbname --package=$package --template=$template \
 --model=$modelPackage --model-path=$modelPath --dao=$daoPackage --dao-path=$daoPath
