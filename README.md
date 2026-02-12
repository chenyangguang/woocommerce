# woocommerce

WooCommerce Go REST API Library - A comprehensive Go client for the WooCommerce REST API v3.

[![GoDoc](https://godoc.org/github.com/chenyangguang/woocommerce?status.svg)](https://godoc.org/github.com/chenyangguang/woocommerce)
[![Go Report Card](https://goreportcard.com/badge/github.com/chenyangguang/woocommerce)](https://goreportcard.com/report/github.com/chenyangguang/woocommerce)

## Features

- Full WooCommerce REST API v3 coverage
- Context support for cancellation and timeout
- Automatic retry with rate limiting
- Comprehensive error handling
- Type-safe API responses

## Installation

```console
go get github.com/chenyangguang/woocommerce
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    woo "github.com/chenyangguang/woocommerce"
)

func main() {
    // Create an app with your WooCommerce credentials
    app := woo.App{
        CustomerKey:    "your_consumer_key",
        CustomerSecret: "your_consumer_secret",
    }

    // Create a client for your shop
    client := app.NewClient("your-shop.com")

    // List products
    var products []woo.Product
    err := client.Product.List(nil, &products)
    if err != nil {
        log.Fatal(err)
    }

    for _, p := range products {
        fmt.Printf("Product: %s - %s\n", p.Name, p.Price)
    }
}
```

## API Coverage

| Resource | Methods |
|----------|---------|
| **Products** | List, Get, Create, Update, Delete, Batch |
| **Product Variations** | List, Get, Create, Update, Delete, Batch |
| **Product Categories** | List, Get, Create, Update, Delete, Batch |
| **Product Tags** | List, Get, Create, Update, Delete, Batch |
| **Product Attributes** | List, Get, Create, Update, Delete, Batch |
| **Product Shipping Classes** | List, Get, Create, Update, Delete, Batch |
| **Product Reviews** | List, Get, Create, Update, Delete, Batch |
| **Orders** | List, Get, Create, Update, Delete, Batch |
| **Order Notes** | List, Get, Create, Delete |
| **Order Refunds** | List, Get, Create, Delete |
| **Customers** | List, Get, Create, Update, Delete, Batch |
| **Coupons** | List, Get, Create, Update, Delete, Batch |
| **Payment Gateways** | List, Get, Update |
| **Webhooks** | List, Get, Create, Update, Delete, Batch |
| **Settings** | Get, Update |

## Context Support

All API methods have context-aware variants for timeout and cancellation control:

```go
import "context"

// Create a context with timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Use context-aware methods
var products []woo.Product
err := client.Product.ListWithContext(ctx, nil, &products)
```

## Error Handling

The library provides typed errors for proper error handling:

```go
import "errors"

err := client.Product.Get(999, nil, &product)
if err != nil {
    var respErr woo.ResponseError
    if errors.As(err, &respErr) {
        fmt.Printf("API Error (Status %d): %s\n", respErr.Status, respErr.Message)
    }

    var rateLimitErr woo.RateLimitError
    if errors.As(err, &rateLimitErr) {
        fmt.Printf("Rate limited. Retry after %d seconds\n", rateLimitErr.RetryAfter)
    }
}
```

## Configuration Options

```go
// With custom HTTP client
client := app.NewClient("your-shop.com", woo.WithHTTPClient(customClient))

// With custom logger
client := app.NewClient("your-shop.com", woo.WithLogger(myLogger))

// With retry support
client := app.NewClient("your-shop.com", woo.WithRetry(3))
```

## Documentation

For complete API documentation, see:
- [GoDoc](https://godoc.org/github.com/chenyangguang/woocommerce)
- [WooCommerce REST API Docs](https://woocommerce.github.io/woocommerce-rest-api-docs/)

## Requirements

- Go 1.21+
- WooCommerce 3.5+

## License

MIT License - see [LICENSE](LICENSE) for details.
