# go-jenius
An Unofficial SDK for Jenius Payment Channel

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


## License

MIT

Copyright (c) 2020 - Indra Octama