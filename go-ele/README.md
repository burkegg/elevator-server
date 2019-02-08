# Golang Version of the Elevator Server
This doc should give you links to everything you need to know to develop
the elevator-server code in Go (currently tested in Go 1.8.x).

# Generate
We are using the [`go-swagger`](github.com/go-swagger/go-swagger/cmd/swagger)
tool to build the generic scaffolding for our project.

First, let's get the repo, and put it in the correct location:

`$GOPATH` should already be defined if you have a working GO install

```
$ mkdir -p $GOPATH/src/github.com/onbeep/
$ cd $GOPATH/src/github.com/onbeep/
$ git clone git@github.com:OnBeep/elevator-server.git
```
At this point we're referencing `$GOPATH/src/github.com/onbeep/elevator-server`
as `$REPO_ROOT`.

Now, install the code generation tool:
```
$ go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

Let's generate the server code:, and run it:
```
$ cd $REPO_ROOT/go-ele
$ swagger generate server -f ../elevator.yml
```

The generated code needs you to install some more libs:
```
$ go get -u github.com/go-openapi/runtime github.com/tylerb/graceful github.com/jessevdk/go-flags golang.org/x/net/context
```

Now build the executable, and run it:
```
$ go install ./...
$ $GOPATH/bin/lifty-server --port 5000
>> serving lifty at http://127.0.0.1:5000
```

You should be able to see the server responding:
```
$ curl http://127.0.0.1:5000/v1/welcome
>> {"msg":"Welcome to the Elevator Server"}
```

To run the unit tests:
```
$ cd $REPO_ROOT/go-ele
$ go test ./...
```
