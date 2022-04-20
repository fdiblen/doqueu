# doqueu - Docker container management API

## Generate API docs

```console
go get -u github.com/fdiblen/doqueu
swag init
```

## Run the server

```console
go run main.go
```

API documentation can be fount at [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)


## Examples

List containers

```console
curl -X 'GET' \
  'http://localhost:8080/api/v1/containers' \
  -H 'accept: application/json'
```

Run a container

```console
curl -X 'POST' \
  'http://localhost:8080/api/v1/containers/run' \
  -H 'accept: application/json' \
  -H 'Content-Type: multipart/form-data' \
  -F 'imagename=ubuntu:latest' \
  -F 'command=echo "Hello World"'
```

Stop a running container

```console
curl -X 'POST' \
  'http://localhost:8080/api/v1/containers/stop' \
  -H 'accept: application/json' \
  -H 'Content-Type: multipart/form-data' \
  -F 'id=15b16c8bae4eea58f77fbf9558cb487502565e7a4bb4d170c0fc93884e532c7c'
```
