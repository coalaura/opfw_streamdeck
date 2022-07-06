@echo off

echo Building...
mkdir dist
go-winres make
go build -o "dist\OP-FW Streamdeck.exe"

echo Done
pause