package main

// Generic types

type identity struct {
	MPinID []byte
	Token  []byte
	DTAs   string
}

type header struct {
	Key, Value string
}

// Registration responses

type sessionResponse struct {
	AccessURL string `json:"accessURL"`
	QRURL     string `json:"qrURL"`
	WebOTT    string `json:"webOTT"`
}

type verificationURLResponse struct {
	VerificationURL string `json:"verificationURL"`
}

type confirmationRequest struct {
	UserID           string `json:"userId"`
	Code             string `json:"code"`
	ConfirmMissmatch bool   `json:"confirmMissmatch"`
}

type confirmationResponse struct {
	ActivateToken           string `json:"actToken"`
	AccessID                string `json:"accessId"`
	State                   string `json:"state"`
	Nonce                   string `json:"nonce"`
	ExpireTime              int64  `json:"expireTime"`
	VerificationRedirectURL string `json:"verificationRedirectUrl"`
}

type registerResponse struct {
	Active     bool   `json:"active"`
	AppID      string `json:"appId"`
	CustomerID string `json:"customerId"`
	ExpireTime int    `json:"expireTime"`
	MPinID     string `json:"mpinId"`
	NowTime    int    `json:"nowTime"`
	RegOTT     string `json:"regOTT"`
}

type signatureResponse struct {
	ClientSecretShare string `json:"clientSecretShare"`
	CS2URL            string `json:"cs2url"`
	Curve             string `json:"curve"`
	DTAs              string `json:"dtas"`
}

type clientSecretResponse struct {
	ClientSecret    string `json:"clientSecret"`
	DVSClientSecret string `json:"dvsClientSecret"`
	CreatedAt       int    `json:"createdAt"`
	Message         string `json:"message"`
	Version         string `json:"version"`
}

// Authentication responses

type pass1Response struct {
	Y string `json:"y"`
}

type pass2Response struct {
	AuthOTT string `json:"authOTT"`
}

type authenticateResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type accessResponse struct {
	Status      string `json:"status"`
	StatusCode  int    `json:"statusCode"`
	UserID      string `json:"userId"`
	RedirectURL string `json:"redirectURL"`
}
