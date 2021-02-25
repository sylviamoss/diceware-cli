GOARCH=$(shell go env GOARCH)

build: 
	@go build -o ~/go/bin/diceware-cli .