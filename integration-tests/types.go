package main

// Generic types

type identity struct {
	MPinID     []byte
	Token      []byte
	DTAs       string
	PrivateKey []byte
	PublicKey  []byte
}

// MPinID + PublicKey.
func (id *identity) dvsMPinID() []byte {
	return append(id.MPinID, id.PublicKey...)
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

type confirmationResponse struct {
	ActivateToken           string `json:"actToken"`
	AccessID                string `json:"accessId"`
	State                   string `json:"state"`
	Nonce                   string `json:"nonce"`
	ExpireTime              int64  `json:"expireTime"`
	VerificationRedirectURL string `json:"verificationRedirectUrl"`
}

type registrationResponse struct {
	MPinID           string   `json:"mpinId"`
	ProjectID        string   `json:"projectId"`
	DTAs             string   `json:"dtas"`
	Curve            string   `json:"curve"`
	SecretURLs       []string `json:"secretUrls"`
	VerificationType string   `json:"verificationType"`
	PINLength        int      `json:"pinLength"`
	MPinIDCreatedAt  int      `json:"createdAt"`
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
