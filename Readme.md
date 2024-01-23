# Project info
This is a pet golang project. 
GRPC server on golang using clean architecture.

## Components
| Component      |   Vendor   |
|----------------|:----------:|
| Database       | PostgreSQL |
| DB driver      |    Sqlx    |
| Configurations |   Viper    |

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


## Frontend

Code generating for grpc client (In backend root dir)

```
protoc $(find ./internal/infrastructure/ports/grpc/proto -iname "*.proto") \
 			--proto_path=./internal/infrastructure/ports/grpc/proto \
			--plugin=protoc-gen-grpc-web=./frontend/node_modules/.bin/protoc-gen-grpc-web \
			--js_out=import_style=commonjs:./frontend/src/proto \
			--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./frontend/src/proto
```