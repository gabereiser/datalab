version:=v0.0.1

ifeq ($(OS),Windows_NT)
	FOLDER:=windows
	ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
		EXECUTE:=datalab-windows-amd64.exe
	endif
	ifeq ($(PROCESSOR_ARCHITECTURE),ARM64)
		EXECUTE:=datalab-windows-arm64.exe
	endif
else
	UNAME_S := $(shell uname -s)
	UNAME_P := $(shell uname -m)
	ifeq ($(UNAME_S),Linux)
		FOLDER:=linux
		ifeq ($(UNAME_P),AMD64)
			EXECUTE:=datalab-linux-amd64
		endif
		ifeq ($(UNAME_P),ARM64)
			EXECUTE:=datalab-linux-arm64
		endif
	endif
	ifeq ($(UNAME_S),Darwin)
		FOLDER:=darwin
		ifeq ($(UNAME_P),AMD64)
			EXECUTE:=datalab-darwin-amd64
		endif
		ifeq ($(UNAME_P),ARM64)
			EXECUTE:=datalab-darwin-arm64
		endif
		ifeq ($(UNAME_P),arm64)
			EXECUTE:=datalab-darwin-arm64
		endif
	endif
endif
BIN:=$(FOLDER)/$(EXECUTE)
PLATFORMS := linux/amd64/ linux/arm64/ windows/amd64/.exe windows/arm64/.exe darwin/amd64 darwin/arm64

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))
ext = $(word 3, $(temp))


all: clean protoc dep compile

compile: $(PLATFORMS)
	@echo "Done"

run: $(PLATFORMS) 
	@./build/$(BIN)

clean:
	@rm -rf rpc && go clean && rm -rf build

dep:
	@go mod tidy

protoc:
	@mkdir -p rpc
	@cd proto && protoc --go_out=../rpc --go_opt=paths=source_relative --go-grpc_out=../rpc --go-grpc_opt=paths=source_relative *.proto

$(PLATFORMS):
	@echo "Building for $(os) $(arch)"
	@go generate
	@GOOS=$(os) GOARCH=$(arch) go build -ldflags="-X 'datalab.Version=$(version)'" -o 'build/$(os)/datalab-$(os)-$(arch)$(ext)' cmd/main.go

.PHONY: clean dep compile $(PLATFORMS)
