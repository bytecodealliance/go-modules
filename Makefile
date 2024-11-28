wit_files = $(sort $(shell find testdata -name '*.wit' ! -name '*.golden.*'))

# json recompiles the JSON intermediate representation test files.
.PHONY: json
json: $(wit_files)

.PHONY: $(wit_files)
$(wit_files): wit/wasm-tools.wasm
	wasmtime --dir testdata wit/wasm-tools.wasm component wit -j --all-features $@ > $@.json

# golden recompiles the .golden.wit test files.
.PHONY: golden
golden: json
	go test ./wit -update

# generated generated writes test Go code to the filesystem
.PHONY: generated
generated: clean json
	go test ./wit/bindgen -write

.PHONY: clean
clean:
	rm -rf ./generated/*

# tests/generated writes generated Go code to the tests directory
.PHONY: tests/generated
tests/generated: json
	go generate ./tests

# build builds the cmd/wit-bindgen-go binary
.PHONY: build
build: wit/wasm-tools.wasm
	go build -o wit-bindgen-go ./cmd/wit-bindgen-go

wit/wasm-tools.wasm:
	cd wasmtools && \
	git submodule update --init --recursive && \
	cargo build --target wasm32-wasip1 --release --no-default-features -F component
	mv wasmtools/target/wasm32-wasip1/release/wasm-tools.wasm wit/wasm-tools.wasm

# test runs Go and TinyGo tests
GOTESTARGS :=
.PHONY: test
test:
	go test $(GOTESTARGS) ./...
	GOARCH=wasm GOOS=wasip1 go test $(GOTESTARGS) ./...
	tinygo test $(GOTESTARGS) ./...
	tinygo test -target=wasip1 $(GOTESTARGS) ./...
	tinygo test -target=wasip2 $(GOTESTARGS) ./...
	tinygo test -target=wasip2 $(GOTESTARGS) ./tests/...
