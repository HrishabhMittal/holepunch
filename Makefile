.PHONY: build run clean

build:
	@echo building...
	@go build -o ./build/main ./cmd/main

run:
	@go run ./cmd/main

clean:
	@echo cleaning...
	@rm -rf ./build
