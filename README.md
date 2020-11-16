# Go Lang Crash Course

Go is an open source programming language that makes it easy to build simple, reliable and efficient software.

- Ease of use and productivity
- Efficiency and static typing
- Networking performance and built for multi-core processing
- Built for high-throughput, concurrent processing in modern distrubuted apps

## Language Features

- Has the same performance as C
- Is much easier to maintain than Java
- Wide adoption: from #20 to #13 on TIOBE index from 2019 to 2020
- Automatic documentation via GoDoc
- Static code analysis via GoMetaLinter
- Embedded testing environment
- Race condition detection
- Integrated package management, compiler, and test suite + runner

## Installation

```sh
# Download the Go tarball from golang.org (Linux intall)
# Extract and add executables to PATH env variable
tar -C /usr/local -xzf go1.15.5.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin # add this to .profile or .bashrc
go version # Check that install succeeded
export GOPATH=$HOME/$(go env GOPATH) # Set Go path (separate from Go executables)
export PATH=$PATH:$(go env GOPATH)/bin # Add Go path executables to PATH
go get github.com/aws/aws-sdk-go/aws # Add a package from remote repo (creates folders in GOPATH)
```

## Code Organization

- Keep all Go code in a single *workspace*
- A workspace contains one or many version controlled *repositories*
- Each repo contains one or more *packages*
- Each package consists of one or more *Go source files* in a single directory
- Path to a package determines its *import path*

Here's a sample workspace:

```sh
bin/
  hello               # command executable
  outyet              # command executable
src/
  github.com/golang/example/
    .git/             # git repository metadata
    hello/
      hello.go        # command source
    outyet/
      main.go         # command source
      main_test.go    # test source
    stringutil/
      reverse.go      # package source
      reverse_test.go # test source
  golang.org/x/image/
    .git/             # git repository metadata
    bmp/
      reader.go       # package source
      writer.go       # package source


```

## Core Go Commands

```sh
go bug # start a bug report
go build # compile packages and dependencies
go clean # remove object and cached files
go doc # show documentation for package or symbol
go env # print Go environment information
go fix # update packages to use new APIs
go fmt # gofmt (reformat) package sources
go generate # generate Go files by processing source
go get # download and install packages and dependencies
go install # compile and install packages and dependencies
go list # list packages or modules
go mod # module maintenance
go run # compile and run Go program
go test # test packages
go version # print Go version
go vet # report likely mistakes in packages
```