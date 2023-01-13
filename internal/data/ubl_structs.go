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

type CreateAccountRequestStruct struct {
	BranchCode             string `json:"BranchCode"`
	BankAcNum              string `json:"BankAcNum"`
	BankAcTitle            string `json:"BankAcTitle"`
	BankShortName          string `json:"BankShortName"`
	Amount                 int    `json:"Amount"`
	TransDate              string `json:"TransDate"`
	TransTime              string `json:"TransTime"`
	CustomerName           string `json:"CustomerName"`
	FatherHusbandName      string `json:"FatherHusbandName"`
	CNIC                   string `json:"CNIC"`
	CNIC_Expiry            string `json:"CNIC_Expiry"`
	Gender                 string `json:"Gender"`
	DateOfBirth            string `json:"DateOfBirth"`
	MailingAddress         string `json:"MailingAddress"`
	Country                string `json:"Country"`
	City                   string `json:"City"`
	Nationality            string `json:"Nationality"`
	ResidentNumber         string `json:"ResidentNumber"`
	OfficeNumber           string `json:"OfficeNumber"`
	MobileNo               string `json:"MobileNo"`
	RegisteredEmailAddress string `json:"RegisteredEmailAddress"`
	Occupation             string `json:"Occupation"`
	AnnualIncome           string `json:"AnnualIncome"`
	Religion               string `json:"Religion"`
	MaritalStatus          string `json:"MaritalStatus"`
	ZakatDeduction         string `json:"ZakatDeduction"`
	PreferredContactNumber string `json:"PreferredContactNumber"`
	PlanCode               string `json:"PlanCode"`
	PlanName               string `json:"PlanName"`
	SendAccountStatement   string `json:"SendAccountStatement"`
	CountryOfBirth         string `json:"CountryOfBirth"`
	Nationality2           string `json:"Nationality2"`
	CountryOfResince       string `json:"CountryOfResince"`
	CountryTaxResidence    string `json:"CountryTaxResidence"`
	BenificialOwner        string `json:"BenificialOwner"`
	USAPerson              string `json:"USAPerson"`
}

type CreateAccountResponseStruct struct {
	ResultStatus     int                                  `json:"resultStatus"`
	ExceptionMessage string                               `json:"exceptionMessage"`
	Value            CreateAccountResponseNestedObjStruct `json:"value"`
}

type CreateAccountResponseNestedObjStruct struct {
	AccountNo       string `json:"accountNo"`
	ReferenceNumber string `json:"referenceNumber"`
}
