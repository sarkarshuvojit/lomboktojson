test:
	go test ./...
testv:
	go test ./... -v
cover:
	go test ./... -cover
dep:
	go mod download

wasm:
	@GOOS=js GOARCH=wasm go build -o docs/assets/lombok2json.wasm cmd/l2j4wasm/main.go
