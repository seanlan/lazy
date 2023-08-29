if [ $# != 2 ] ; then
    echo "Usage: $0 <model> <api-name>"
    exit 1
fi
model=$1
api=$2
apiout="internal/api/v1"
modelout="internal/model"
serviceout="internal/service"
package="github.com/seanlan/lazy"
template=""
lazy api --package=$package --template=$template \
  --model=$model --api=$api --api-out=$apiout --model-out=$modelout --service-out=$serviceout