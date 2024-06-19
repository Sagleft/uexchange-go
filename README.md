![logo](logo.jpg)

Library for working with the exchange UUSD and Crypton written in Golang

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GoDoc](https://godoc.org/github.com/sagleft/uexchange-go?status.svg)](https://godoc.org/gopkg.in/sagleft/uexchange-go.v1)
[![go-report](https://goreportcard.com/badge/github.com/Sagleft/uexchange-go)](https://goreportcard.com/report/github.com/Sagleft/uexchange-go)

## What do you get by using this library:
* crypto currency converter api;
* crypton crypto currency price api;
* cryptocurrency trading api;
* crypto exchange and wallet;
* anonymous cryptocurrency exchange api.

[CRP.is](https://crp.is/) - is the best anonymous bitcoin exchange, on the basis of which you can make your own bot using this library in the Golang language.

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

## Useful resources

* [Getting access to the Utopia API](https://udocs.gitbook.io/utopia-api/utopia-api/how-to-enable-api-access).
* [Examples of projects](https://udocs.gitbook.io/utopia-api/utopia-api/examples-of-projects) created based on the Utopia API.
* [How to create WEB 3.0 chatbots](https://udocs.gitbook.io/utopia-api/utopia-api/creating-chat-bots).
* [How to earn in WEB 3.0](https://udocs.gitbook.io/utopia-api/how-to-earn-in-web-3.0)
* [Official website](https://u.is/en/) of the Utopia P2P project.
