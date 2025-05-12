test:
	go test ./...
testv:
	go test ./... -v
cover:
	go test ./... -cover
dep:
	go mod download
