#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(dirname "$0")"
ROOT_DIR="$SCRIPT_DIR/.."

cd "$ROOT_DIR" || exit

ENV_FILE=".env"
ENV_CONTENT="ENV=dev"

# Check if the file does not exist, then create it
if [ ! -f "$ENV_FILE" ]; then
    echo -e "$ENV_CONTENT" > "$ENV_FILE"
    echo "$ENV_FILE created."
else
    echo "$ENV_FILE already exists."
fi

# Run dev
air