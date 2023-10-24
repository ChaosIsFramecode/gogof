# Project makefile

all: run

OUTPUT = ./bin/GameOfLife

ifeq ($(OS),Windows_NT)
OUTPUT = ./bin/GameOfLife.exe
endif

run:
	go run main.go
build:
	go build -o $(OUTPUT) main.go
