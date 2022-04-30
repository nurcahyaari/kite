generate:
	go generate .

build: generate
	mkdir build && go build -o build/app

test: generate
	go test ./...