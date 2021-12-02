model=$1
api=$2
apiout="app/api/v1"
modelout="app/model"
serviceout="app/service"
package="github.com/seanlan/lazy"
template=""
lazy api --package=$package --template=$template \
  --model=$model --api=$api --api-out=$apiout --model-out=$modelout --service-out=$serviceout