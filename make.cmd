setlocal
call :"%1"
endlocal
exit /b

:""
    set GOARCH=386
    go fmt
    go build -ldflags "-s -w"
    exit /b

:"package"
    for %%I in (%CD%) do set "NAME=%%~nI"
    zip "%NAME%-windows-386-%DATE:/=%" %NAME%.exe
    exit /b
