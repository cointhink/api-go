all: main

main: src/cointhink/**/*go
	GOPATH=`pwd` go build -o cointhink src/cointhink/cmd/*.go

install:
	GOPATH=`pwd` go get github.com/elgs/gojq
	GOPATH=`pwd` go get github.com/golang/protobuf/jsonpb
	GOPATH=`pwd` go get github.com/golang/protobuf/proto
	GOPATH=`pwd` go get github.com/golang/protobuf/ptypes
	GOPATH=`pwd` go get github.com/golang/protobuf/ptypes/any
	GOPATH=`pwd` go get github.com/hjson/hjson-go
	GOPATH=`pwd` go get github.com/jmoiron/sqlx
	GOPATH=`pwd` go get github.com/ogier/pflag
	GOPATH=`pwd` go get github.com/satori/go.uuid
	GOPATH=`pwd` go get github.com/google/uuid
	GOPATH=`pwd` go get gopkg.in/gomail.v2
	GOPATH=`pwd` go get github.com/gorilla/websocket
	GOPATH=`pwd` go get github.com/lib/pq

protoc3:
	wget https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip
	mkdir protoc3
	unzip protoc-3.3.0-linux-x86_64.zip -d protoc3
	rm protoc-3.3.0-linux-x86_64.zip

watch:
	while true; do echo; inotifywait -r src/cointhink -e MODIFY 2> /dev/null; gofmt -w .; make main; done
