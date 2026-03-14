# lzt-api-go

Go API wrapper for Lolzteam Forum & Market. Auto-generated from OpenAPI schemas.

## Installation

```bash
go get github.com/fantolla123/lzt-api-go
```

## Quick Start

```go
package main

import (
	"fmt"

	lztapi "github.com/fantolla123/lzt-api-go"
	"github.com/fantolla123/lzt-api-go/forum"
	"github.com/fantolla123/lzt-api-go/market"
)

func main() {
	f := forum.NewClient("your_token")
	threads, err := f.ThreadsList()
	if err != nil {
		panic(err)
	}
	fmt.Println(threads)

	m := market.NewClient("your_token")
	accounts, err := m.CategoryList()
	if err != nil {
		panic(err)
	}
	fmt.Println(accounts)
}
```

## Proxy

```go
import "time"

f := forum.NewClient("your_token",
	lztapi.WithProxy("http://user:pass@proxy:8080"),
	lztapi.WithTimeout(60 * time.Second),
)
```

## Error Handling

```go
import "errors"

result, err := f.ThreadsGet(123)
if err != nil {
	var rateLimitErr *lztapi.RateLimitError
	var authErr *lztapi.AuthError
	var notFoundErr *lztapi.NotFoundError

	switch {
	case errors.As(err, &rateLimitErr):
		fmt.Println("Rate limited")
	case errors.As(err, &authErr):
		fmt.Println("Invalid token")
	case errors.As(err, &notFoundErr):
		fmt.Println("Not found")
	}
}
```

## Code Generation

```bash
go run ./cmd/codegen --schema schemas/forum.json --output forum/ --package forum
go run ./cmd/codegen --schema schemas/market.json --output market/ --package market
```

## Features

- Auto-retry on 429/502/503 with exponential backoff
- Proxy support
- Typed response models
- Multipart file upload support
- No external dependencies (stdlib only)
- 151 forum methods, 115 market methods

## License

MIT
