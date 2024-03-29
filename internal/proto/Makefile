proto_gen_export_path:
	export PATH=$PATH:$(go env GOPATH)/bin

generate_messageBox_proto:
	 protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       messageBox.proto