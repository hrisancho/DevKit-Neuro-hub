protobuf:
	protoc --proto_path=./proto/ --go_out=./proto/ --go_opt=paths=source_relative ./proto/*.proto
submodule:
	git submodule update --remote --merge
