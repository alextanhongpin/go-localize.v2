generate:
	@go generate ./...

deps:
	@go get -u golang.org/x/tools
	@go install golang.org/x/text/cmd/gotext@latest
