build-gofherd:
	cd gofherd && GOARCH=amd64 GOOS=linux go build -o gofherd-bin && cd ..

deploy: rebuild
	eval $(docker-machine env gofherd)
	docker-compose up

rebuild: build-gofherd
	docker-compose build gofherd caddy grafana prometheus
