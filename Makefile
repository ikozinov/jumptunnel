run:
	set -a; source .env; set +a; go run cmd/main.go 

tidy:
	go mod tidy
