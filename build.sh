#!/bin/bash

GOOS=darwin GOARCH=386 go build -o chow_Darwin_386 main.go models.go

GOOS=darwin GOARCH=amd64 go build -o chow_Darwin_x86_64 main.go models.go

GOOS=linux GOARCH=amd64 go build -o chow_Linux_x86_64 main.go models.go

GOOS=linux GOARCH=386 go build -o chow_Linux_386 main.go models.go