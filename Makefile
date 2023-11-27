include .env

.SILENT:

build:
	protoc --proto_path=./internal/infrastructure/ports/grpc/proto \
 		--go_out=./internal/infrastructure/ports/grpc/proto_gen/ \
 		--go-grpc_out=./internal/infrastructure/ports/grpc/proto_gen/ \
 		--grpc-gateway_out=./internal/infrastructure/ports/grpc/proto_gen/ \
 		--openapiv2_out ./internal/infrastructure/ports/http/api/ \
 		--openapiv2_opt ignore_comments=true \
 		--openapiv2_opt allow_merge=true \
 		--openapiv2_opt allow_merge=true \
 		--openapiv2_opt generate_unbound_methods=false \
 		./internal/infrastructure/ports/grpc/proto/*.proto
