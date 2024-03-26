#!/usr/bin/env bash

GOOS=linux GOARCH=arm go build -o wol_arm wol.go
GOOS=linux GOARCH=amd64 go build -o wol_amd64 wol.go