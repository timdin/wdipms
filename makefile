run:
	docker compose up --remove-orphans

build:
	docker compose build

update:
	make build
	make run
