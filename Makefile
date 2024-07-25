build_linux_amd64:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o biker main.go
