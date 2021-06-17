#!/bin/bash

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/linux/seo-tools .
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows/seo-tools.exe .
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/mac/seo-tools .