# go-server

test server in golang

## DESCRIPTION

The go-server is for testing server in golang. When you do http request to a specific endpoint, it will be returned is http response.

## INSTALLATION

```
$ go get github.com/Ken2mer/go-server
```

## USAGE

Run the following.

```
$ make run
```

From another console, run the following.

```
$ curl http://127.0.0.1:8080/hello/YourName
hello, YourName!
```

### use gin-gonic/gin

```
$ make test
$ go run gin/server
...
[GIN-debug] GET    /user/:name               --> main.setupRouter.func1 (3 handlers)
[GIN-debug] GET    /user/:name/*action       --> main.setupRouter.func2 (3 handlers)
[GIN-debug] GET    /data                     --> main.setupRouter.func3 (3 handlers)
[GIN-debug] GET    /json                     --> main.setupRouter.func4 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```

