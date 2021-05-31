build:
	go build -o httpd main.go

dev:
	export GIN_MODE=debug
	go run httpd/main.go

prod:
	export GIN_MODE=release
	go run httpd/main.go