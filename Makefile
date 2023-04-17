run-httpserver:
	go run ./cmd/httpserver/...

run-migration:
	go run ./cmd/migration/...

start-mysql:
	$(MAKE) stop-mysql
	docker run \
		-d \
		--rm \
		--name productmanagement-mysql \
		-p 3307:3306 \
		--env-file .env \
		mysql:5.7

stop-mysql:
	-docker kill \
		productmanagement-mysql