# go-github #

go-nodeping is a minimal Go client library for accessing the [Nodeping API][].

**Documentation:** [![GoDoc](https://godoc.org/github.com/jswank/nodeping-go/nodeping?status.svg)](https://godoc.org/github.com/jswank/go-nodeping/nodeping)  

go-nodeping requires Go version 1.4 or greater.

## Usage ##

```go
import "github.com/jswank/go-nodeping/nodeping"
```

Construct a new Nodeping client, then use the various services on the client to
access different parts of the Nodeping API.

### Authentication ###

The go-nodeping library does not directly handle authentication. Instead, when
creating a new client, pass an `http.Client` that can handle authentication for
you. 

```go
import "golang.org/x/oauth2"

func main() {

  api_token := "your api token"

	tp := nodeping.BasicAuthTransport{
		Token: strings.TrimSpace(api_token),
	}

	client := nodeping.NewClient(tp.Client())

  // list all checks
	checks, _, err := client.Checks.List()

}
```

## License ##

This library is distributed under the BSD-style license found in the [LICENSE](./LICENSE)
file.
