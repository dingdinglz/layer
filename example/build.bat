@echo off
set GOARCH=wasm
set GOOS=js
go build -o web/app.wasm
::请自行替换你的平台和架构
set GOARCH=amd64
set GOOS=windows
go build
