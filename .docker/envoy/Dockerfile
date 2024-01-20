#FROM    alpine:latest AS build
#WORKDIR /app
#RUN     apk add --no-cache protoc && \
#        apk add protobuf-dev
#COPY    ./internal/infrastructure/ports/grpc/proto .
#
## build proto descriptor
#RUN     protoc --include_imports --include_source_info \
#            --proto_path=./ \
#            --descriptor_set_out=new.pb new.proto
#
#FROM    envoyproxy/envoy:v1.22.0
#COPY    --from=build /app/new.pb /tmp/
#RUN chmod go+r /etc/envoy/envoy.yaml


FROM    alpine:3.15.1 AS build

WORKDIR /app

RUN     apk add --no-cache protoc && \
        apk add protobuf-dev

COPY    ./internal/infrastructure/ports/grpc/proto .

# build proto descriptor
RUN     protoc --include_imports --include_source_info \
            --proto_path=./ \
            --descriptor_set_out=new.pb new.proto

FROM    envoyproxy/envoy:v1.22.0

COPY    --from=build /app/new.pb /tmp/
