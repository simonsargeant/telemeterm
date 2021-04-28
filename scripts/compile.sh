#!/bin/bash
set -eufo pipefail
IFS=$'\t\n'

# Compile the binary to run locally

go build -o ./bin/telemeterm ./cmd/telemeterm
