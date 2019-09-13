# Go Workshop for Beginners

## Getting Started

### Installing GO 

#### Home/Linuxbrew

It may be convenient to install the latest version of Go through the
[Homebrew](https://brew.sh/) and [Linuxbrew](http://linuxbrew.sh/) package
managers.

```
$ brew install go
```

#### Install with Binary Distributions

The [https://golang.org/dl/](https://golang.org/dl/) page contains distros for
Windows, MacOS, Linux, and source. The
[installation instructions](https://golang.org/doc/install) explains how to
install them.


#### Verify your Installation

```bash
$ go version   
go version go1.13 darwin/amd64
```

### Getting the Workshop Code

```
$ go get github.com/solidnerd/go-workshop
```

or 

```
$ mkdir -p $GOPATH/src/github.com/solidnerd/go-workshop
$ git clone $GOPATH/src/github.com/solidnerd/go-workshop
```

Check your installation with a first running Go Programm

```bash
$ go run github.com/solidnerd/go-workshop
```


Aftwards switch to the Workshop Dir

```bash
$ cd $GOPATH/src/github.com/solidnerd/go-workshop
```


## Resources 

A lot of inspiration and examples comes from here [sendwithus/workshop-go](https://github.com/sendwithus/workshop-go)
