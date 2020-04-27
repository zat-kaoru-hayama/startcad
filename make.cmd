setlocal
set GOARCH=386
go fmt
go build -ldflags "-s -w"
endlocal
