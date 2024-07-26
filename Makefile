.PHONY: run build
run:
	@ go run main.go
build:
	@ go build -o bq-schema main.go
