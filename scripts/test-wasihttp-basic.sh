#!/bin/bash
set -euo pipefail

echo "Building basic example..."
cd ./x/wasihttp/examples
tinygo build -target=wasip2-http.json -o basic.wasm ./basic

echo "Running basic example with wasmtime..."
wasmtime serve --addr 0.0.0.0:8088 -Scli basic.wasm &
SERVER_PID=$!

echo "Testing / endpoint..."
ROOT_RESPONSE=$(curl -s 'http://0.0.0.0:8088/')
if [[ "$ROOT_RESPONSE" != "Hello world!"* ]]; then
	echo "ERROR: Root endpoint test failed"
	echo "Expected response containing 'Hello world!'"
	echo "Actual response: $ROOT_RESPONSE"
	kill $SERVER_PID 2>/dev/null
	exit 1
fi

kill $SERVER_PID 2>/dev/null
rm basic.wasm

echo "All basic tests passed!"
