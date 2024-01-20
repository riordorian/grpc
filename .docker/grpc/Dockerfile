FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

#ENTRYPOINT CompileDaemon --build="go build cmd/main.go" --command=./runserver
ENTRYPOINT /go/bin/CompileDaemon --build "go build cmd/main.go" -command=./main
