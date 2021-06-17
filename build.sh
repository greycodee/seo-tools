#!/bin/bash

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/seo-tools-mac/seo-tools .
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/seo-tools-win/seo-tools.exe .
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/seo-tools-linux/seo-tools .