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