package jenius

import (
	"bytes"
	"encoding/json"
	"github.com/indraoct/go-jenius/helper"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Client struct {
	UrlDevelopment  string   // Provide by Jenius
	UrlProduction   string   // Provide by Jenius
	Environtment    string   // dev or prod
	ClientId        string   // Provide by Jenius
	ApiKey          string   // Provide by Jenius
	ApiSecret       string   // Provide By Jenius
	ChannelId       string   // Provide By Jenius
	CallbackUrl     string   // Your Callback Url
	
	TimeStamp       string   // Define by yourself example : helper.GetNowTime().Format(jenius.JENIUS_DATETIME_FORMAT)
	ReferenceNo     string   // Your Transaction ID
	
	Debug 			bool            // To Show Log in your bash / shell
	Logger 			*log.Logger
}

func NewClient() Client{

	return Client{
		Environtment: "",
		Debug        : true,
		Logger: log.New(os.Stderr, "", log.LstdFlags),
	}
}

var timeout = 30 * time.Second
var httpClient = http.Client{Timeout: timeout}

func (c *Client) GenerateTokenAccess(payloadString string,httpVerb string,relativeUrl string)(credentials CredentialAccess,err error){
	
	var (
		req                 *http.Request
		baseUrl             string
		client              http.Client
		respToken           ResponseAccessToken
		StringtoSignature   string
	)
	
	if c.Environtment == "dev"{
		baseUrl = c.UrlDevelopment
	}else if c.Environtment == "prod"{
		baseUrl = c.UrlProduction
	}
	
	param := url.Values{}
	param.Set("grant_type","client_credentials")
	
	payload := bytes.NewBufferString(param.Encode())
	req, err = http.NewRequest("POST", baseUrl+":8089/api/oauth/token", payload)
	
	if err != nil{
		if c.Debug == true{
			c.Logger.Println("Request generate token failed: ", err)
		}
		return credentials,err
	}
	
	req.Header.Add("Authorization","Basic "+Base64Encode(c.ClientId))
	req.Header.Add("Content-Type","application/x-www-form-urlencoded")
	
	res, err := client.Do(req)
	if err != nil {
		if c.Debug == true{
			c.Logger.Println("Request generate token failed: ", err)
		}
		return credentials,err
	}
	defer res.Body.Close()
	
	err = json.NewDecoder(res.Body).Decode(&respToken)
	if err != nil {
		if c.Debug == true{
			c.Logger.Println("Request generate token failed: ", err)
		}
		return credentials,err
	}
	
	if httpVerb == "POST" {
		StringtoSignature = strings.Replace(httpVerb+":"+relativeUrl+":"+c.ApiKey+":"+
			c.TimeStamp+":"+payloadString, " ", "", -1)
	}else if httpVerb == "GET"{
		StringtoSignature = strings.Replace(httpVerb+":"+relativeUrl+":"+c.ApiKey+":"+
			c.TimeStamp, " ", "", -1)
	}else{
		StringtoSignature = strings.Replace(httpVerb+":"+relativeUrl+":"+c.ApiKey+":"+
			c.TimeStamp, " ", "", -1)
	}
	
	credentials.BTPNSignature         = ComputeHmac256(StringtoSignature,c.ApiSecret)
	credentials.Authorization         = respToken.AccessToken
	credentials.BTPNApiKey            = c.ApiKey
	credentials.BTPNTimeStamp         = c.TimeStamp
	credentials.XTransmissionDateTime = c.TimeStamp
	
	return credentials,err
}


func (c *Client) ExecuteRequest(req *http.Request) ([]byte, error) {
	if c.Debug {
		c.Logger.Println("Request ", req.Method, ": ", req.URL.Host + req.URL.Path)
	}
	
	start := time.Now()
	res, err := httpClient.Do(req)
	
	if c.Debug {
		c.Logger.Println("Completed in ", time.Since(start))
		c.Logger.Println("Transaction Date", helper.GetNowTime().Format(JENIUS_DATETIME_FORMAT))
	}
	
	if err != nil {
		if c.Debug {
			c.Logger.Println("Request failed: ", err)
		}
		return nil, err
	}
	
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		if c.Debug {
			c.Logger.Println("Cannot read response body: ", err)
		}
		return nil, err
	}
	
	if c.Debug {
		c.Logger.Println("API response: ", string(resBody))
	}
	
	return resBody, nil
}