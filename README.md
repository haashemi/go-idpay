# Go [IDPay](https://idpay.ir/) API Wrapper

[![Go Report Card](https://goreportcard.com/badge/github.com/haashemi/go-idpay)](https://goreportcard.com/report/github.com/haashemi/go-idpay) [![Go Reference](https://pkg.go.dev/badge/github.com/haashemi/go-idpay.svg)](https://pkg.go.dev/github.com/haashemi/go-idpay)


### Installation
```
go get github.com/haashemi/go-idpay
```

------

## Usage

```go
package main

import (
    "fmt"
    "github.com/haashemi/go-idpay"
)

func main() {
    // 1- create new idpay client
    idp := idpay.New("My-IDPay-APIKey")

    // 2- create new transaction
    tr, _ := idp.CreateTransaction(orderID, callbackURL, amount, nil)
    
    // 3- verify the transaction
    trInfo, _ := idp.VerifyTransaction(tr.ID, orderID)
}
```
