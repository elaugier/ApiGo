set GOOS=
set GOARCH=
go build ./...
go install ./...
go build -o dist\apigo-engine\win-apigo-engine.exe cmd\apigo-engine\main.go
go build -o dist\apigo-worker\win-apigo-worker.exe cmd\apigo-worker\main.go
REM set GOOS=darwin 
REM set GOARCH=amd64 
REM go build dist\apigo-engine\mac-apigo-engine cmd\apigo-engine\main.go
REM go build dist\apigo-worker\mac-apigo-worker cmd\apigo-worker\main.go
REM set GOOS=linux 
REM set GOARCH=amd64 
REM go build dist\apigo-engine\lnx-apigo-engine cmd\apigo-engine\main.go
REM go build dist\apigo-worker\lnx-apigo-worker cmd\apigo-worker\main.go
REM set GOOS=
REM set GOARCH=