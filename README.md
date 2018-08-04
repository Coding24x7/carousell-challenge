# carousell-challenge

## Prerequisite

### Setup Go dev environment using instruction at 

https://golang.org/doc/install


### Setup dependency management tool

```bash
go get github.com/tools/godep
```

#### Get all the go code dependencies in vendor folder
```bash
godep save
```

### For Rest API and swagger code generation by goagen

```bash
go get github.com/goadesign/goa
./goagen.sh
```

## Step to run

```bash
./run.sh
```

The API documentation is available at `localhost:8080/`.

Swagger documentation provides facilities to call the REST API as well, so it can be used for testing.

## Workflow

This is simple implementation of twitter cloning. It has two main endpoints user and topics. 
To work on topics we need to have user available.

### User

It has three endpoints

#### Register user
It will create new user account.

#### Delete user
It will remove user account which is already registered with service.

#### Login user account
It will login to user account by using input username and password.


### Topic

It has three endpoints

#### Post topic
It will post new topic for input user with given content.

#### Show topic
It will show top 20 topics sorted by upvotes (in descending order).

#### Upvote/Downvote topic
It will add users' vote for given topic.


### Testing

```bash
go test ./lib/
```