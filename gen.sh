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

protoc-go-inject-tag -input=./pokemon.pb.go  
(cd breed-move-service && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)
(cd breed-service && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)
(cd gateway-api && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)
(cd move-service && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)
(cd pokemon-service && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)
(cd battle-service && gv update github.com/Mrcampbell/pgo2/protorepo/pokemon)

################################



gv update github.com/Mrcampbell/pgo2/protorepo/pokemon
