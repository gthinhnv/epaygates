#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(dirname "$0")/.."
SRC_DIR="$ROOT_DIR/../services/metadata-service/proto"
DEST_DIR="$ROOT_DIR/proto"

mkdir -p "$DEST_DIR"

echo "Syncing all proto folders from $SRC_DIR (except sharedpb) to $DEST_DIR..."

for DIR in "$SRC_DIR"/*; do
    # Skip non-directories
    [ -d "$DIR" ] || continue

    BASENAME="$(basename "$DIR")"

    # Skip sharedpb
    if [[ "$BASENAME" == "sharedpb" ]]; then
        continue
    fi

    DEST_SUBDIR="$DEST_DIR/$BASENAME"

    # Ensure destination subdir exists
    mkdir -p "$DEST_SUBDIR"

    # Copy all files AND folders inside DIR (but not DIR itself)
    cp -r "$DIR/"* "$DEST_SUBDIR/" 2>/dev/null || true

    # Make proto files read-only
    find "$DEST_SUBDIR" -type f -name "*.proto" -exec chmod -w {} +

    echo " → Synced folder: $BASENAME"
done

echo "Done syncing metadata-service proto folders."
