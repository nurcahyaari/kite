
dev: generate
	go run kite.go wire_gen.go ${action} --name ${name} --path ${path}

create_new:
	go run kite.go wire_gen.go new --name ${name} --path ${path}

create_domain:
	go run kite.go wire_gen.go domain --name ${name} --path ${path} ${create-only-folder}

generate:
	go generate .

build: generate
	mkdir build && go build -o build/app

test: generate
	go test ./...