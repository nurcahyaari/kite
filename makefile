generate:
	go generate .

build: generate
	mkdir -p build && go build -o build/kite

test: generate
	go test ./...