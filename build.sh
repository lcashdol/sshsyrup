#!/bin/bash
echo "Compiling binary.."
go build -ldflags "-s -w" -o sshsyrup ./cmd/syrup
