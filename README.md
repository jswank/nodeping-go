# go-github #

go-nodeping is a minimal Go client library for accessing the [Nodeping API][].

**Documentation:** [![GoDoc](https://godoc.org/github.com/jswank/nodeping-go/nodeping?status.svg)](https://godoc.org/github.com/jswank/go-nodeping/nodeping)  

go-nodeping requires Go version 1.4 or greater.

## Usage ##

```go
import "github.com/jswank/go-nodeping/nodeping"

func main() {

  client := nodeping.NewClient("Nodeping API token")

  // list all checks
  checks, _, err := client.Checks.List()

}
```

Construct a new Nodeping client, then use the various services on the client to
access different parts of the Nodeping API.

## License ##

This library is distributed under the BSD-style license found in the [LICENSE](./LICENSE)
file.
