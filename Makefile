include .env

.SILENT:

build:
	protoc --proto_path=./internal/infrastructure/ports/grpc/proto \
		--go_opt=paths=source_relative \
 		--go_out=./internal/infrastructure/ports/grpc/proto_gen/ \
 		--go-grpc_out=./internal/infrastructure/ports/grpc/proto_gen/ \
 		--grpc-gateway_out=./internal/infrastructure/ports/grpc/proto_gen/ \
 		--grpc-gateway_opt generate_unbound_methods=true \
 		--openapiv2_out \
 		./internal/infrastructure/ports/grpc/proto/*.proto
