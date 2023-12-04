build: tidy
	go build -ldflags "-s -w" -o bin/ .

run: tidy
	go run main.go

compile: tidy icon
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/iconize-linux-amd64 .
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o dist/iconize-windows-amd64.exe .
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o dist/iconize-macos-amd64 .
	GOOS=freebsd GOARCH=amd64 go build -ldflags "-s -w" -o dist/iconize-freebsd-amd64 .

tidy:
	go mod tidy

icon:
	go install github.com/akavel/rsrc@latest
	rsrc -ico go.ico

loc:
	go install github.com/boyter/scc/v3@latest
	scc --exclude-dir vendor --exclude-dir bin .
