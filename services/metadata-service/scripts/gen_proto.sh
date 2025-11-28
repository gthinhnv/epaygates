#!/usr/bin/env bash
set -e

ROOT_DIR=$(dirname "$0")/..
PROTO_ROOT="$ROOT_DIR/proto"
OUT_DIR="$ROOT_DIR/gen/go"

echo "Cleaning old generated files..."
rm -rf "$OUT_DIR"
mkdir -p "$OUT_DIR"

echo "Generating proto files..."
find "$PROTO_ROOT" -name "*.proto" | while read -r proto_file; do
    echo "Processing $proto_file ..."
    protoc \
        -I "$PROTO_ROOT" \
        --go_out="$OUT_DIR" \
        --go_opt=paths=source_relative \
        --go-grpc_out="$OUT_DIR" \
        --go-grpc_opt=paths=source_relative \
        "$proto_file"
done

echo "Done!"
