run:
	docker compose up

build:
	docker compose build

update:
	make build
	make run
