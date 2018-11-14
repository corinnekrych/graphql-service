M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: build
build: vendor/ schema ; $(info $(M) Building project...)
	go build ./...

.PHONY: clean
clean: ; $(info $(M) [TODO] Removing generated files... )
	$(RM) schema/bindata.go
	$(RM) -rf vendor/

.PHONY: schema
schema: schema/bindata.go ; $(info $(M) Embedding schema files into binary...)

schema/bindata.go: vendor/ ./schema/*.graphql ./schema/types/*.graphql
	PATH=$(GOPATH)/bin:$(PATH) go generate ./schema

vendor/: Gopkg.toml Gopkg.lock ; $(info $(M) Fetching dependencies...)
	go get github.com/golang/dep/cmd/dep
	go get -u github.com/jteeuwen/go-bindata/...
	dep ensure

.PHONY: server
server: schema ; $(info $(M) Starting development server...)
	go run main.go
