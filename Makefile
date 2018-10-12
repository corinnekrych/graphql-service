M = $(shell printf "\033[34;1m▶\033[0m")

build: dep ; $(info $(M) Building project...)
	go build

clean: ; $(info $(M) [TODO] Removing generated files... )
	$(RM) schema/bindata.go

dep: setup ; $(info $(M) Ensuring vendored dependencies are up-to-date...)
	dep ensure

schema: dep ; $(info $(M) Embedding schema files into binary...)
	PATH=$(GOPATH)/bin:$(PATH) go generate ./schema

setup: ; $(info $(M) Fetching github.com/golang/dep...)
	go get github.com/golang/dep/cmd/dep
	go get -u github.com/jteeuwen/go-bindata/...

server: schema ; $(info $(M) Starting development server...)
	go run main.go

.PHONY: build clean container dep image schema setup server
