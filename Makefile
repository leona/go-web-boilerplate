.PHONY: default
default: run-dev

build:
	@go generate ./bin
	@go build -o ./bin/run ./bin/server.go
	
build-small:
	@go generate ./bin
	@go build -ldflags="-s -w" -o ./bin/run ./bin/server.go
	@upx --brute ./bin/run
	#sudo apt-get install upx-ucl
	
run-dev:
	@go generate ./bin
	@go run ./bin/server.go
	
run:
	@./bin/run
	
setup:
	@export GOPATH=$HOME/workspace
	@export PATH=$PATH:$GOPATH/bin
	
install: 
	@go get -u github.com/kataras/iris/iris
	@go get -u github.com/iris-contrib/middleware/secure
	@go get -u github.com/iris-contrib/middleware/logger
	@go get -u github.com/valyala/quicktemplate/qtc
	@go get -u github.com/jaredfolkins/badactor
	@go get github.com/fatih/color
	@go get github.com/boltdb/bolt