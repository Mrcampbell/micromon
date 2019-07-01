protoc --go_out=plugins=grpc:. *.proto

protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  **/*.proto

protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  **/*.proto

protoc-go-inject-tag -input=./pokemon/pokemon.pb.go  

gv update github.com/Mrcampbell/pgo2/protorepo/pokemon

protoc-go-inject-tag -input=./pokemon.pb.go  
(cd breed-move-service && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)
(cd breed-service && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)
(cd gateway-api && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)
(cd move-service && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)
(cd pokemon-service && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)

NLU=nsqlookupd
NLU_C="$NLU | tee 1.log | sed -e 's/^/[nsqlookupd] /'"

NSQD="nsqd --lookupd-tcp-address=127.0.0.1:4160 -broadcast-address=127.0.0.1"
NSQD_C="$NSQD | tee 2.log | sed -e 's/^/[nsqd] /'"

NSQADMIN="nsqadmin --lookupd-http-address=127.0.0.1:4161"
NSQADMIN_C="$NSQADMIN | tee 3.log | sed -e 's/^/[nsqadmin] /'"

$NLU_C & $NSQD_C & $NSQADMIN_C
