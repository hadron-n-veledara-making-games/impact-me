.PHONY: build
build:
	go build -o telegramreceiver.exe -v ./cmd/telegrambot
	go build -o telegramworker.exe -v ./cmd/telegramworker