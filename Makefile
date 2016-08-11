.PHONY: default
default: run-dev

build:
	@go generate ./bin
	@go build -o ./bin/run ./bin/server.go
	
run-dev:
	@go generate ./bin
	@go run ./bin/server.go
	
run:
	@./bin/run
	
setup:
	@mkdir $HOME/workspace
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