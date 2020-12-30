CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../targets/bgm-linux-amd64 ../main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ../targets/bgm-macos-amd64 ../main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ../targets/bgm-windows-x64.exe ../main.go