
<p align="center"><img width="320px" src="https://upload.wikimedia.org/wikipedia/id/8/89/Jenius-logo.png">  <img height="100px" src="https://stickershop.line-scdn.net/stickershop/v1/product/1349132/LINEStorePC/main.png;compress=true"></p>

# go-jenius
Implementation Jenius Pay using Go (Golang)

![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/indraoct/go-jenius) ![GitHub](https://img.shields.io/github/license/indraoct/go-jenius)

## 1. Installation
To install this library , you can import it into your GOPATH by running this command in your terminal :
```text
go get github.com/indraoct/go-jenius
```

## 2. Usage

To use this library , first thing first is you need to initiate the client property data like this :

```go
	
	client := jenius.NewClient()
    	client.UrlDevelopment = "{JENIUS URL DEVELOPMENT}"
    	client.UrlProduction  = "{JENIUS URL PRODUCTION}"
    	client.Environtment = "dev"     // dev for development ; prod for production
    	client.ApiKey       = "{YOUR API KEY}"
    	client.ApiSecret    = "{YOUR API SECRET}"
    	client.ChannelId    = "{YOUR CHANNEL ID}"
    	client.ClientId     = "{YOUR CLIENT ID}"
    	client.CallbackUrl  = "{YOUR CALLBACK URL}"
    	client.ReferenceNo  = "{YOUR TRANSACTION ID}"
    	client.TimeStamp    = helper.GetNowTime().Format(jenius.JENIUS_DATETIME_FORMAT) //timestamp standard for jenius payment
    	
```

And then you must parse the initiate data into function initiate like this :

```go
   //parse the client data into function
   	function := jenius.Function{
   		Client:client,
   	}
```

And then you can set the parameter data for main function , for example "PayRequest" for creating payment
After that you can call the main function such as Create Transaction and Payment Status.

## 3. Example

### Create Transaction

Create Transaction is a jenius service to send transaction data from client to jenius server.

```go
package main

import (
	"fmt"
	"github.com/indraoct/go-jenius"
	"github.com/indraoct/go-jenius/helper"
)

func main(){
	
	//initiate client
	client := jenius.NewClient()
	client.UrlDevelopment = "{JENIUS URL DEVELOPMENT}"
	client.UrlProduction  = "{JENIUS URL PRODUCTION}"
	client.Environtment = "dev"     // dev for development ; prod for production
	client.ApiKey       = "{YOUR API KEY}"
	client.ApiSecret    = "{YOUR API SECRET}"
	client.ChannelId    = "{YOUR CHANNEL ID}"
	client.ClientId     = "{YOUR CLIENT ID}"
	client.CallbackUrl  = "{YOUR CALLBACK URL}"
	client.ReferenceNo  = "{YOUR TRANSACTION ID}"
	client.TimeStamp    = helper.GetNowTime().Format(jenius.JENIUS_DATETIME_FORMAT) //timestamp standard for jenius payment
	
	//match client into function
	function := jenius.Function{
		Client:client,
	}
	
	//PayRequest Parameter
	req := &jenius.PayRequest{
		Cashtag:"{YOUR CUSTOMER CASHTAG}",
		PromoCode:"",
		PurchaseDesc:"Beli Emas 0.6 gr Indra",
		TxnAmount:"350000",
	}
	
	//main function
	res,err := function.CreateTransaction(req)
	
	if err != nil{
		fmt.Println(err.Error())
	}
	
	fmt.Println(res)
}
```

### Transaction Status

Transaction status is a jenius service to get data transaction status.

```go
package main

import (
	"fmt"
	"github.com/indraoct/go-jenius"
	"github.com/indraoct/go-jenius/helper"
)

func main(){
	
	//initiate client
	client := jenius.NewClient()
	client.UrlDevelopment = "{JENIUS URL DEVELOPMENT}"
	client.UrlProduction  = "{JENIUS URL PRODUCTION}"
	client.Environtment = "dev"     // dev for development ; prod for production
	client.ApiKey       = "{YOUR API KEY}"
	client.ApiSecret    = "{YOUR API SECRET}"
	client.ChannelId    = "{YOUR CHANNEL ID}"
	client.ClientId     = "{YOUR CLIENT ID}"
	client.CallbackUrl  = "{YOUR CALLBACK URL}"
	client.TimeStamp    = helper.GetNowTime().Format(jenius.JENIUS_DATETIME_FORMAT) //timestamp standard for jenius payment
	
	//match client into function
	function := jenius.Function{
		Client:client,
	}
	
	//Status Request Parameter
	req := &jenius.TrxStatusRequest{
		ReferenceNo:"{YOUR TRANSACTION ID}",
		TrxDateTime:"2020-04-08T15:16:45.791+07:00", // transaction date time in jenius format
	}
	
	//main function
	res,err := function.TransactionStatus(req)
	
	if err != nil{
		fmt.Println(err.Error())
	}
	
	fmt.Println(res)
}
```

## License

MIT

Copyright (c) 2020 - Indra Octama