#!/usr/bin/env bash

gox -os="darwin " -arch="amd64" -output=./bin/ini-op ./...

# "./bin/ini-op.{{.OS}}.{{.Arch}}"
