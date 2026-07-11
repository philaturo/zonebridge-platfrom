.PHONY: setup dev backend frontend docker stop test lint build clean

setup:
	task setup

dev:
	task dev

backend:
	task backend

frontend:
	task frontend

docker:
	task docker

stop:
	task stop

test:
	task test

lint:
	task lint

build:
	task build

clean:
	task clean