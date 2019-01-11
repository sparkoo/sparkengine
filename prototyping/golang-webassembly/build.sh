#!/usr/bin/env bash

set -x

rm -rf out/
mkdir -p out

cp index.html out/
cp wasm_exec.js out/

GOOS=js GOARCH=wasm go build -o out/main.wasm

go run server/server.go
