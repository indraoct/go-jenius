package jenius

type PayRequest struct {
	TxnAmount       string  `json:"txn_amount"`
	Cashtag         string  `json:"cashtag"`
	PromoCode       string  `json:"promo_code"`
	UrlCallback     string  `json:"url_callback"`
	PurchaseDesc    string  `json:"purchase_desc"`
}

type CredentialAccess struct {
	Authorization         string  `json:"authorization"`
	BTPNSignature         string  `json:"btpn_signature"`
	BTPNApiKey            string  `json:"btpn_api_key"`
	BTPNTimeStamp         string  `json:"btpn_time_stamp"`
	XChannelId            string  `json:"x_channel_id"`
	XNode                 string  `json:"x_node"`
	XTransmissionDateTime string  `json:"x_transmission_date_time"`
	XReferenceNo          string  `json:"x_reference_no"`
}
