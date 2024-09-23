main_package_path = ./cmd/gowc
binary_name = gowc


## tidy: tidy modfiles and format .go files
.PHONY: tidy
tidy:
	go mod tidy -v
	go fmt ./...

## build: build the application
.PHONY: build
build:
	go build -o=/tmp/bin/${binary_name} ${main_package_path}

## test: run all tests
.PHONY: test
test:
	go test -v ./...
