package jenius

type ResponseAccessToken struct {
	AccessToken     string     `json:"access_token"`
	TokenType       string     `json:"token_type"`
	ExpiresIn       int        `json:"expires_in"`
	Scope           string     `json:"scope"`
}


type PayResponseSuccess struct {
	ResponseCode    string      `json:"response_code"`
	ResponseDesc    struct{
		Id      string      `json:"id"`
		En      string      `json:"en"`
	}      `json:"response_desc"`
}


type PayResponseError struct {
	ErrorCode    string      `json:"ErrorCode"`
	ErrorMessage    struct{
		Indonesian      string      `json:"Indonesian"`
		English       string       `json:"English"`
	}      `json:"ErrorMessage"`
}


type GeneralResponsePayment struct {
	Approval        string         `json:"approval"`
	ResponseCode    string      `json:"response_code"`
	ResponseDesc    struct{
		Id      string      `json:"id"`
		En      string      `json:"en"`
	}      `json:"response_desc"`
	ErrorCode    string      `json:"ErrorCode"`
	ErrorMessage    struct{
		Indonesian      string      `json:"Indonesian"`
		English       string       `json:"English"`
	}      `json:"ErrorMessage"`
}

type JeniusCallback struct {
	Approval                string         `json:"approval"`
	ResponseCode            string         `json:"response_code"`
	ResponseDesc            struct{
		Id          string      `json:"id"`
		En          string      `json:"en"`
	}  `json:"response_desc"`
}