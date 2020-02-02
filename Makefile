GOARCH=$(shell go env GOARCH)

build: 
	@go build -o ~/go/bin/diceware .

build_windows:
	@GOOS=windows go build -o ./pkg/diceware . \
		&& zip ./pkg/diceware_windows_$(GOARCH).zip ./pkg/diceware \
		&& rm ./pkg/diceware

build_linux:
	@GOOS=linux go build -o ./pkg/diceware . \
		&& zip ./pkg/diceware_linux_$(GOARCH).zip ./pkg/diceware \
		&& rm ./pkg/diceware

build_macOS:
	@go build -o ./pkg/diceware . \
		&& zip ./pkg/diceware_macOS_$(GOARCH).zip ./pkg/diceware \
		&& rm ./pkg/diceware