install:
proto-build:
	protoc \
		--proto_path=. \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$$(find proto -name "*.proto")

build:
	go build .
build-server:
	go build . GOOS=linux GOARCH=amd64

test:
	go test -race ./...
test-coverage:
lint:
	golangci-lint run ./...
lint-fix:
	golangci-lint run --fix ./...
infra-up:
infra-down:
docker-up:
run:
