#!/bin/bash

echo "FOOOOOOOOO"

### build linux

go build -v -o build/f1-telemetry main.go

### build windows

GOOS=windows GOARCH=386 CXX=i686-w64-mingw32-g++ CC=i686-w64-mingw32-gcc go build -v -ldflags -H=windowsgui -o build/f1-telemetry.exe main.go