run: build
	go run ./cmd/server/main.go

build:
	GOOS=js GOARCH=wasm go build -o ./assets/main.wasm ./cmd/wasm/main.go 
