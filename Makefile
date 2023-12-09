build:
	go mod tidy
	go build -ldflags "-s -w" -o bin/ .

run:
	go mod tidy
	go run .

icon:
	go install github.com/akavel/rsrc@latest
	rsrc -ico go.ico

windows: icon build
