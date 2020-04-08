package jenius

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Function struct {
	Client
}

// function create transaction
func (f *Function) CreateTransaction(payReq *PayRequest) (res GeneralResponsePayment, err error) {
	
	var (
		 credentials    CredentialAccess
		 path           string
		 httpVerb       string
		 relativeUrl    string
		 payLoad        []byte
		 payLoadString  string
		 jsonPayReq     []byte
		 req            *http.Request
		)
	
	//initiate value
	httpVerb             = "POST"
	relativeUrl          = "/jpay/payrequest"
	payLoad,_            = json.Marshal(payReq)
	payLoadString        = string(payLoad)
	credentials,err      = f.Client.GenerateTokenAccess(payLoadString,httpVerb,relativeUrl)
	jsonPayReq, _ 	     = json.Marshal(payReq)
	path                 = EndpointCreateTransaction
	
	//check endpoint environtment
	if f.Client.Environtment == "dev"{
		path = f.Client.UrlDevelopment + path
	}else if f.Client.Environtment == "prod"{
		path = f.Client.UrlProduction + path
	}else{
		f.Client.Logger.Println("Parameter Environtment must be 'dev' or 'prod' ")
		return res, errors.New("Parameter Environtment must be 'dev' or 'prod' ")
	}
	
	//create req http request
	req,err = http.NewRequest(httpVerb,path, bytes.NewBuffer(jsonPayReq))
	
	if err != nil{
		if f.Client.Debug == true{
			f.Client.Logger.Println("Request creation failed: ", err)
		}
	}
	
	req.Header.Add("Authorization","Bearer "+credentials.Authorization)
	req.Header.Add("BTPN-Signature",credentials.BTPNSignature)
	req.Header.Add("BTPN-ApiKey",credentials.BTPNApiKey)
	req.Header.Add("BTPN-Timestamp",f.Client.TimeStamp)
	req.Header.Add("X-Channel-Id",f.Client.ChannelId)
	req.Header.Add("X-Node","Jenius Pay")
	req.Header.Add("X-Transmission-Date-Time",f.Client.TimeStamp)
	req.Header.Add("X-Reference-No",f.Client.ReferenceNo)
	req.Header.Add("Content-Type","application/json")
	
	//execute http request
	resp, err := f.Client.ExecuteRequest(req)
	if err != nil{
		f.Client.Logger.Println("Request creation failed: ", err)
		return res,err
	}
	
	err = json.Unmarshal(resp,&res)
	if err != nil{
		f.Client.Logger.Println("Request creation failed: ", err)
		return res,err
	}
	
	return res, nil
}


//function transaction status
func (f *Function) TransactionStatus(statusReq *TrxStatusRequest) (res PayResponseSuccess, err error) {
	
	var (
		credentials    CredentialAccess
		path           string
		httpVerb       string
		relativeUrl    string
		payLoad        []byte
		payLoadString  string
		req            *http.Request
	)
	
	//initiate value
	httpVerb             = "GET"
	relativeUrl          = "/jpay/paystatus"
	payLoad,_            = json.Marshal(statusReq)
	payLoadString        = string(payLoad)
	credentials,err      = f.Client.GenerateTokenAccess(payLoadString,httpVerb,relativeUrl)
	path                 = EndpointTransactionStatus
	
	//check endpoint environtment
	if f.Client.Environtment == "dev"{
		path = f.Client.UrlDevelopment + path
	}else if f.Client.Environtment == "prod"{
		path = f.Client.UrlProduction + path
	}else{
		f.Client.Logger.Println("Parameter Environtment must be 'dev' or 'prod' ")
		return res, errors.New("Parameter Environtment must be 'dev' or 'prod' ")
	}
	
	//create req http request
	req,err = http.NewRequest(httpVerb,path, bytes.NewBuffer([]byte{}))
	if err != nil{
		if f.Client.Debug == true{
			f.Client.Logger.Println("Request creation failed: ", err)
		}
	}
	
	req.Header.Add("Authorization","Bearer "+credentials.Authorization)
	req.Header.Add("BTPN-Signature",credentials.BTPNSignature)
	req.Header.Add("BTPN-ApiKey",credentials.BTPNApiKey)
	req.Header.Add("BTPN-Timestamp",f.Client.TimeStamp)
	req.Header.Add("X-Channel-Id",f.Client.ChannelId)
	req.Header.Add("X-Node","Jenius Pay")
	req.Header.Add("X-Transmission-Date-Time",f.Client.TimeStamp)
	req.Header.Add("X-Original-Transmission-Date-Time",statusReq.TrxDateTime)
	req.Header.Add("X-Reference-No",statusReq.ReferenceNo)
	req.Header.Add("cache-control","no-cache")
	
	//execute http request
	resp, err := f.Client.ExecuteRequest(req)
	if err != nil{
		f.Client.Logger.Println("Request creation failed: ", err)
		return res,err
	}
	
	err = json.Unmarshal(resp,&res)
	if err != nil{
		f.Client.Logger.Println("Request creation failed: ", err)
		return res,err
	}
	
	return res,nil
}
