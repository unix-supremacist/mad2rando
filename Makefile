default:
	go build -o bin/mad2rando mad2rando.go
	CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o bin/mad2rando.exe mad2rando.go

run:
	go run mad2rando.go