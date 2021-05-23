compile:
	echo "Compiling for all platforms"
	GOOS=darwin GOARCH=arm64 go build -o bin/darwin-arm64/fix-fb-meta *.go
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin-amd64/fix-fb-meta *.go
	GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/fix-fb-meta *.go
	GOOS=windows GOARCH=386 go build -o bin/win-x86/fix-fb-meta.exe *.go
	GOOS=windows GOARCH=amd64 go build -o bin/win-x64/fix-fb-meta.exe *.go

run:
	go run *.go