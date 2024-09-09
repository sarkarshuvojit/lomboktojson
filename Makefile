test:
	@go test ./...
testv:
	@go test ./... -v
dep:
	@go mod download
wasm:
	@GOOS=js GOARCH=wasm go build -o ui/main.wasm 
