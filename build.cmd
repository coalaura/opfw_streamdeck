@echo off

echo Building...
mkdir dist

SET CGO_ENABLED=1

go-winres make
go build -ldflags "-H=windowsgui" -o "dist\OP-FW Streamdeck.exe"

go build -o "dist\debug.exe"

echo Done
pause