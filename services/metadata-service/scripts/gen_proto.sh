#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR=$(dirname "$0")/..

echo "Running buf generate..."
buf dep update
buf generate
echo "Done!"
