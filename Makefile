version:=v0.0.1

all: clean dep compile

compile: build
	@echo "Done"

run: build
	@./bin/datalab

clean:
	@rm -rf rpc && go clean && rm -rf bin && rm -rf public && cd ui && rm -rf node_modules

dep:
	@go mod tidy

build:
	@echo "Building"
	@go generate
	@go build -ldflags="-X 'github.com/gabereiser/datalab.Version=$(version)'" -o 'bin/datalab' cmd/main.go

.PHONY: clean dep compile build
