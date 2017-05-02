all: main

main: src/*go
	go build -o main src/*.go

protoc3:
	wget https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip
	mkdir protoc3
	unzip protoc-3.3.0-linux-x86_64.zip -d protoc3
	rm protoc-3.3.0-linux-x86_64.zip

watch:
	while true; do inotifywait -r src -e MODIFY; make main; done
