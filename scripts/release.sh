#!/usr/bin/env bash

REPO_DIR="$(dirname "$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )")"

go build goapprove.go && mv goapprove $REPO_DIR/bin/goapprove