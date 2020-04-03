package jenius

import (
	"bytes"
	"encoding/json"
	"errors"
)

type Function struct {
	Client
}

func (f *Function) CreateTransaction(req *PayRequest) (res GeneralResponsePayment, err error) {
	
	var credentials CredentialAccess
	
	httpVerb        := "POST"
	relativeUrl     := "/jpay/payrequest"
	payLoad,_       := json.Marshal(req)
	payLoadString   := string(payLoad)
	
	credentials,err  = f.Client.GenerateTokenAccess(payLoadString,httpVerb,relativeUrl)
	
	jsonReq, _ 	    := json.Marshal(req)
	
	body, err := f.Call("POST", EndpointCreateTransaction, bytes.NewBuffer(jsonReq),credentials)
	if err != nil {
		f.Client.Logger.Println("Error buy init: ", err)
		return res, err
	}
	
	json.Unmarshal(body, &res)
	if err != nil {
		f.Client.Logger.Println(err.Error())
		return res, errors.New(err.Error())
	}
	
	return res, nil
}