#!/usr/bin/env bash
OUTPUT=$(revive -config config/revive.toml ./...)
if [ -z "$OUTPUT" ]; then
    echo "No errors"
else
    echo "$OUTPUT"
fi