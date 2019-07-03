protoc --go_out=plugins=grpc:. *.proto

#############################


cd protorepo 

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

cd ../

(cd breed-move-service && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)
(cd breed-service && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)
(cd gateway-api && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)
(cd move-service && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)
(cd pokemon-service && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)
(cd battle-service && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)

################################


#############################


cd protorepo 

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

cd ../

(cd breed-move-service && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)
(cd breed-move-service && gv remove +u)

(cd breed-service && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)
(cd breed-service && gv remove +u)

(cd gateway-api && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)
(cd gateway-api && gv remove +u)

(cd move-service && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)
(cd move-service && gv remove +u)

(cd pokemon-service && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)
(cd pokemon-service && gv remove +u)

(cd battle-service && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)
(cd battle-service && gv gv remove +u)

################################


gv update github.com/Mrcampbell/pgo2/protorepo/pokemon
