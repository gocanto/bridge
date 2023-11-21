.PHONY: otnacog\:webhooks

APP_PATH = $(shell pwd)

include .env

run:
	go run main.go