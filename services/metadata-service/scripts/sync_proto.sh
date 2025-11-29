#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR=$(dirname "$0")/..
SRC_DIR="$ROOT_DIR/../../shared/proto/sharedpb"
DEST_DIR="$ROOT_DIR/proto/sharedpb"

mkdir -p "$DEST_DIR"

# Only these files will be synced
FILES=(
	"seo.proto"
	"status.proto"
	"page_type.proto"
	"ads_platform.proto"
)

echo "Syncing selected proto files..."

for FILE in "${FILES[@]}"; do
	SRC_FILE="$SRC_DIR/$FILE"
	DEST_FILE="$DEST_DIR/$FILE"

	if [ ! -f "$SRC_FILE" ]; then
		echo "Warning: $SRC_FILE does not exist, skipping"
		continue
	fi

	# Copy the file
	cp -f "$SRC_FILE" "$DEST_FILE"

	# Make it read-only
	chmod -w "$DEST_FILE"

	echo " → Synced: $FILE"
done

echo "Done syncing proto files."
