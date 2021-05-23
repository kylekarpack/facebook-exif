compile:
	echo "Compiling for all platforms"
	GOOS=darwin GOARCH=arm64 go build -o bin/fix-fb-meta-arm64 *.go
	GOOS=darwin GOARCH=amd64 go build -o bin/fix-fb-meta-amd64 *.go
	GOOS=linux GOARCH=amd64 go build -o bin/fix-fb-meta-amd64-linux *.go
	GOOS=windows GOARCH=386 go build -o bin/fix-fb-meta-x64.exe *.go
	GOOS=windows GOARCH=amd64 go build -o bin/fix-fb-meta-x86.exe *.go

run:
	go run *.go