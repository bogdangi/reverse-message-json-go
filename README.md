# Introduction

Simple HTTP Service reverse message from message service by env variable MESSAGE_SERVICE_URL:

Given env variable `MESSAGE_SERVICE_URL=http://host.docker.internal:3000/` exported
And message service returns
```
{
  "id": "1",
  "message": "Hello world"
}
```
And reverse message service runs on `http://localhost:3001`
When I get `http://localhost:3001/`
Then it returns
```
{
  "id": "1",
  "message": "Hello world"
}
```

# Run tests

```
docker run -it -v $PWD:/go/src/app -w /go/src/app -e CGO_ENABLED=0 golang:alpine go test .
```
