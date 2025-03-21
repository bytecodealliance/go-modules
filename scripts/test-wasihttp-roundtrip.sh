#!/bin/bash
set -euo pipefail

echo "Building roundtrip example..."
cd ./x/wasihttp/examples
tinygo build -target=wasip2-roundtrip.json -o roundtrip.wasm ./roundtrip

echo "Running roundtrip example..."
output=$(wasmtime run -Shttp -Sinherit-network -Sinherit-env roundtrip.wasm)

# Verify GET request worked
if ! echo "$output" | grep -q "Status: 200" || ! echo "$output" | grep -q "https://postman-echo.com/get"; then
	echo "ERROR: GET request verification failed"
	echo "$output"
	exit 1
fi

# Verify POST request worked
if ! echo "$output" | grep -q '"foo": "bar"'; then
	echo "ERROR: POST request verification failed"
	echo "$output"
	exit 1
fi

# Verify PUT request worked
if ! echo "$output" | grep -q '"baz": "blah"'; then
	echo "ERROR: PUT request verification failed"
	echo "$output"
	exit 1
fi

rm roundtrip.wasm
echo "All roundtrip tests passed!"
