build-gofherd:
	cd gofherd && GOARCH=amd64 GOOS=linux go build -o gofherd-bin && cd ..

deploy: build-gofherd
	eval $(docker-machine env gofherd)
	docker-compose up

rebuild: build-gofherd
	docker-compose build gofherd caddy grafana prometheus
