
dev: generate
	go run kite.go wire_gen.go ${action} --name ${name} --path ${path}

run:
	go run kite.go wire_gen.go ${action} --name ${name} --path ${path}

generate:
	go generate .

build: generate
	mkdir build && go build -o build/app
