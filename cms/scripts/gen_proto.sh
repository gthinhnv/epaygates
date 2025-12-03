#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(dirname "$0")"
ROOT_DIR="$SCRIPT_DIR/.."
GEN_DIR="$ROOT_DIR/gen/go"

echo "→ Cleaning $GEN_DIR ..."
if [[ -d "$GEN_DIR" ]]; then
  rm -rf "$GEN_DIR"
fi

echo "→ Running buf generate..."
buf generate

echo "Done!"
