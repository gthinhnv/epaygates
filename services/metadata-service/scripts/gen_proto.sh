#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(dirname "$0")"
GEN_DIR="$SCRIPT_DIR/../gen"

echo "→ Removing generated folder: $GEN_DIR"
rm -rf "$GEN_DIR"
mkdir -p "$GEN_DIR"

echo "→ Syncing proto files..."
bash "$SCRIPT_DIR/sync_proto.sh"

echo "→ Running buf generate..."
buf dep update
buf generate

echo "Done!"
