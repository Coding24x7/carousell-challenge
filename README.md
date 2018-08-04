# carousell-challenge

## Prerequisite

### Setup Go dev environment using instruction at 

https://golang.org/doc/install


### Setup dependency management tool

https://golang.github.io/dep/

#### Get all the go code dependencies in vendor folder
```bash
$ dep ensure -v
```

### For Rest API and swagger code generation by goagen

```bash
go get github.com/goadesign/goa
./goagen.sh
```

## Step to run

```
./run.sh
```

The API documentation is available at `localhost:8080/swagger`.

Swagger documentation provides facilities to call the REST API as well, so it can be used for testing.
