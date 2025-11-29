#!/usr/bin/env bash
set -euo pipefail

# Get the directory of this script
SCRIPT_DIR=$(dirname "$0")

# Run the sync_shared_proto.sh script
bash "$SCRIPT_DIR/sync_shared_proto.sh"
