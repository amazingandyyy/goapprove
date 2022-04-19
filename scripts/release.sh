#!/usr/bin/env bash

export REPO_DIR="$(dirname "$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)")"
export package_name=$(basename $REPO_DIR)

(
	set -x
	go build $REPO_DIR/cmd/main.go
)
mv $REPO_DIR/main $REPO_DIR/bin/$package_name
