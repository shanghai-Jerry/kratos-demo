
API_PROTO_FILES=$(find ./api -name *.proto)
API_PROTO_GEN_FILES=$(find ./api -name *.pb.go)


go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest
go install github.com/favadi/protoc-go-inject-tag@latest

mkdir -p ./api/docs && protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
				 --go-http_out=paths=source_relative:./api \
				 --doc_out=./api/docs \
				 --doc_opt=html,api.html,source_relative \
	       ${API_PROTO_FILES}

# cd proto && protoc-go-inject-tag -input="./*.pb.go"


