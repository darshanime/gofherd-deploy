## gofherd deploy

1. Create a docker machine:

```
docker-machine create --driver digitalocean --digitalocean-access-token=XX --digitalocean-image=ubuntu-18-04-x64 --digitalocean-private-networking=true --digitalocean-region=blr1 --digitalocean-size=s-2vcpu-2gb gofherd
```

2. Edit code in `gofherd/main.go`

3. Run `eval $(docker-machine env gofherd)` and then `make deploy`

4. Visit `:3000` for Grafana dashboard, `:9090` for Prometheus dashboard, `:5555` for Gofherd. Get the machine IP using `docker-machine ip gofherd`
