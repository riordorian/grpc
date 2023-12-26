# Project info

## Migrations
#### Installation
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

#### Usage
https://dev.to/techschoolguru/how-to-write-run-database-migration-in-golang-5h6g

###### Migration create example
```
migrate create -ext sql -dir ./ -seq *MIGRATION_NAME*
```    

###### Migration up example
```
migrate -path ./ -database "postgresql://grpc:password@localhost:5432/grpc?sslmode=disable" -verbose up
```    


## API GW
Krakend used as api gw.

Command example for compiling pb file for krakend
```
protoc -I=internal/infrastructure/ports/grpc/proto --descriptor_set_out=new.pb --include_imports internal/infrastructure/ports/grpc/proto/*.proto
```