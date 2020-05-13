test: docker_build_test
	docker-compose up -d
	docker-compose exec -T http make tests
#	docker-compose stop

pact_consumer: docker_build_test
	docker-compose up -d
	docker-compose exec -T http go test ./consumer
	docker-compose stop

unit_test:
	go test `go list ./... | grep -v e2e_test`

tests:
	go test `go list ./... | grep -v consumer`
#	go test -v ./pact

docker_build:
	docker build . -t template

docker_build_test:
	docker build . -t template_test --target=test

docker_run:
	docker run --publish 8080:8080 template