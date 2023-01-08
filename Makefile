hello:
	echo "Hello"

build:
	@go build -o bin/bedbo

run: build
	@./bin/bedbo

dev: 
	LOG_LEVEL=1 APP_ENV="development" air

debug:
	LOG_LEVEL=1 APP_ENV="development" air -d

test:
	@go test -v ./...