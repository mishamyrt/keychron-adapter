APP_NAME = Keychron Adapter.app
BIN_PATH = build/keychron_adapter

.PHONY: build
build:
	go build -o "$(BIN_PATH)"
	rm -rf "build/$(APP_NAME)"
	cp -r "$(APP_NAME)" "build/"
	mv "$(BIN_PATH)" "build/$(APP_NAME)/Contents/MacOS/"

run:
	go run cli/main.go