GOARCH=$(shell go env GOARCH)

build: 
	@go build -o ~/go/bin/diceware-cli .

build_windows:
	@env GOOS=windows go build . \
		&& zip ./pkg/diceware_windows_$(GOARCH).zip diceware-cli.exe \
		&& rm diceware-cli.exe

build_linux:
	@env GOOS=linux go build . \
		&& zip ./pkg/diceware_linux_$(GOARCH).zip diceware-cli \
		&& rm diceware-cli

build_mac:
	@go build . \
		&& zip ./pkg/diceware_macOS_$(GOARCH).zip diceware-cli \
		&& rm diceware-cli

release: build_mac build_linux build_windows