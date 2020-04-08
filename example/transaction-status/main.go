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