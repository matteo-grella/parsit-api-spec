package api

//go:generate openapi2proto -spec spec.yaml -out service.proto -annotate
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative service.proto
//go:generate protoc --proto_path=. --include_imports --include_source_info --descriptor_set_out=api_descriptor.pb service.proto
//go:generate protoc -I . --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative service.proto

