generate:
	go generate .

build: generate
	mkdir -p build && go build -o build/app

test: generate
	go test ./...