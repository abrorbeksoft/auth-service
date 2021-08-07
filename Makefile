CURRENT_DIR=$(shell pwd)

start:
	  go run cmd/main.go

compile:
		./scripts/generate.sh ${CURRENT_DIR}

.PHONY:compile