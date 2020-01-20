GOARCH=$(shell go env GOARCH)

build: 
	@go build -o ~/go/bin/diceware .

build_windows:
	@GOOS=windows go build -o ./pkg/diceware_windows_$(GOARCH) .

build_linux:
	@GOOS=linux go build -o ./pkg/diceware_linux_$(GOARCH) .
