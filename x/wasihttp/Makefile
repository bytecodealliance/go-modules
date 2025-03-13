.PHONY: tools
tools:
	go generate -tags tools ./...

.PHONY: go-bindings
go-bindings:
	go run go.bytecodealliance.org/cmd/wit-bindgen-go generate -o internal/ ./wit
