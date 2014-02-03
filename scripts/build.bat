@echo off
set APP=runeos
set ADDR=":8808"

set RUNEOS_ROOT=E:\projects\goprj\runeos
set GOPATH=E:\projects\goprj\runeos;E:\projects\goprj\gopath

if exist bin\%APP%.exe del bin\%APP%.exe

echo Build %APP%...
go build -ldflags="-s -w" -o bin\%APP%.exe .\src\.
echo ===================================

bin\%APP% %1 %2 %3 %4 %5 %6
@echo on
