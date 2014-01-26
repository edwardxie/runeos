@echo off
set APP=runeos
set ADDR=":8808"

set GOPATH=E:\projects\goprj\runeos;E:\projects\goprj\gopath

::GO_INSTALL_LIST=utils log4g kernel fuse network network/web network/websocket service/chatroom
::echo Install package %APP%...
::go install %APP%/cmd

if exist bin\%APP%.exe del bin\%APP%.exe

echo Build %APP%...
go build -ldflags="-s -w" -o bin\%APP%.exe .\src\.
::echo GOPATH: %GOPATH%
echo ===================================

bin\%APP% %1 %2 %3 %4 %5 %6
@echo on
