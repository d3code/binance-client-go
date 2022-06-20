# Binance API Client

A Go client for the [Binance API](https://binance-docs.github.io/apidocs/spot/en/#introduction)

### Install

```shell
go get github.com/d3code/binance-client-go
```

### Import

```go
import (
    "github.com/d3code/binance-client-go/binance"
)
```

## Client

To use the client, create a Binance Client.

```go
package main

import (
    "github.com/d3code/binance-client-go/binance"
)

func main() {
    client := binance.Client("<your-api-key>", "<your-api-secret>")
}
```

### Logging

There is no logging in this library, an `error` is returned if there are request / unmarshalling errors and can be handled accordingly.

To print the raw response, call `PrintResponse(true)` on the client.

```go
binance.Client("").PrintResponse(true)
```