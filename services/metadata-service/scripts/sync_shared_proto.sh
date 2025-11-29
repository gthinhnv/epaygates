#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR=$(dirname "$0")/..
SRC_DIR="$ROOT_DIR/../../shared/proto/sharedpb"
DEST_DIR="$ROOT_DIR/proto/sharedpb"

mkdir -p "$DEST_DIR"

echo "Syncing all proto files from $SRC_DIR to $DEST_DIR..."

# Loop through all .proto files in source directory
for SRC_FILE in "$SRC_DIR"/*.proto; do
    # Skip if no files match
    [ -e "$SRC_FILE" ] || continue

    FILE=$(basename "$SRC_FILE")
    DEST_FILE="$DEST_DIR/$FILE"

    # Copy the file
    cp -f "$SRC_FILE" "$DEST_FILE"

    # Make it read-only
    chmod -w "$DEST_FILE"

    echo " → Synced: $FILE"
done

echo "Done syncing shared proto files."
