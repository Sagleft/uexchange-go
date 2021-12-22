![logo](logo.png)

Library for working with the exchange UUSD and Crypton written in Golang

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GoDoc](https://godoc.org/github.com/sagleft/uexchange-go?status.svg)](https://godoc.org/gopkg.in/sagleft/uexchange-go.v1)
[![go-report](https://goreportcard.com/badge/github.com/Sagleft/uexchange-go)](https://goreportcard.com/report/github.com/Sagleft/uexchange-go)

### :link: [Crypton Exchange API docs](https://crp.is/api-doc/)

Install
-----

```bash
go get github.com/Sagleft/uexchange-go
```

```go
import (
	uexchange "github.com/Sagleft/uexchange-go"
)
```

Usage
-----

```go
// create client
client := uexchange.NewClient()

// auth
_, err := client.Auth(uexchange.Credentials{
    AccountPublicKey: "32AE83EF83637ADDA5800E2C9EEB3D456753B0B2CD11D37B90DFA1A1592ED952",
    Password: "mypassword",
})
if err != nil {
    log.Fatalln(err)
}

// get balance
balanceData, err := client.GetBalance()
if err != nil {
    log.Fatalln(err)
}
log.Println(balanceData)

```
