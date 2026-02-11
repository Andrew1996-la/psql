include .env
export

STEP ?= 1

service-run:
	go run main.go

migration-up:
	@migrate -path migration -database ${CONN_STRING} up $(STEP)

migration-down:
	@migrate -path migration -database ${CONN_STRING} down $(STEP)

help:
	@echo "make service-run"
	@echo "make migration-up STEP=2"
	@echo "make migration-down STEP=1"
	@echo "make migration-force VERSION=3"