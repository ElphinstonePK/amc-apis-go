package data

type RequestTokenStruct struct {
	Username string `json:"Username" form:"Username"`
	Password string `json:"Password" form:"Password"`
}

type GetTokenStruct struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"TokenType"`
	ExpiresIn   int    `json:"ExpiresIn"`
	UserName    string `json:"UserName"`
	Issued      string `json:"Issued"`
	Expires     string `json:"Expires"`
}
