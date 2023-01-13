package data

type BalanceInquiryRequestStruct struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	IdNo     string `json:"IdNo"`
	IdType   string `json:"IdType"`
	Stan     string `json:"Stan"`
}

// type BalancesStruct struct {
// 	Mnumenic           string  `json:"mnumenic"`
// 	FUNDID             string  `json:"FUND_ID"`
// 	AGENTID            string  `json:"AGENT_ID"`
// 	AGENTNAME          string  `json:"AGENT_NAME"`
// 	Value              float64 `json:"Value"`
// 	Units              float64 `json:"Units"`
// 	Cost               float64 `json:"Cost"`
// 	RealizedGain       float64 `json:"RealizedGain"`
// 	NavValue           float64 `json:"nav_value"`
// 	NavDate            string  `json:"nav_date"`
// 	FYTD               float64 `json:"FYTD"`
// 	PrevFYTD           float64 `json:"PrevFYTD"`
// 	SinceInceptionGain float64 `json:"SinceInceptionGain"`
// 	InProcessUnits     float64 `json:"InProcessUnits"`
// }

// type BalanceInquiryResponseStruct struct {
// 	PortfolioID        string           `json:"Portfolio_ID"`
// 	CustomerName       string           `json:"Customer_Name"`
// 	Cost               float64          `json:"Cost"`
// 	Value              float64          `json:"Value"`
// 	FYTD               float64          `json:"FYTD"`
// 	PrevFYTD           float64          `json:"PrevFYTD"`
// 	SinceInceptionGain float64          `json:"SinceInceptionGain"`
// 	Balances           []BalancesStruct `json:"Balances"`
// }

type BalanceInquiryResponseStruct []struct {
	PortfolioID        string  `json:"Portfolio_ID"`
	CustomerName       string  `json:"Customer_Name"`
	Cost               float64 `json:"Cost"`
	Value              float64 `json:"Value"`
	FYTD               float64 `json:"FYTD"`
	PrevFYTD           float64 `json:"PrevFYTD"`
	SinceInceptionGain float64 `json:"SinceInceptionGain"`
	Balances           []struct {
		Mnumenic           string  `json:"mnumenic"`
		FUNDID             string  `json:"FUND_ID"`
		AGENTID            string  `json:"AGENT_ID"`
		AGENTNAME          string  `json:"AGENT_NAME"`
		Value              float64 `json:"Value"`
		Units              float64 `json:"Units"`
		Cost               float64 `json:"Cost"`
		RealizedGain       float64 `json:"RealizedGain"`
		NavValue           float64 `json:"nav_value"`
		NavDate            string  `json:"nav_date"`
		FYTD               float64 `json:"FYTD"`
		PrevFYTD           float64 `json:"PrevFYTD"`
		SinceInceptionGain float64 `json:"SinceInceptionGain"`
		InProcessUnits     float64 `json:"InProcessUnits"`
	} `json:"Balances"`
}

type CreateMeezanAccountRequestStruct struct {
	KeyValue struct {
		BasicInfo struct {
			Username          string `json:"Username"`
			FullName          string `json:"FullName"`
			FatherHusbandName string `json:"FatherHusbandName"`
			MotherMaidenName  string `json:"MotherMaidenName"`
			CnicIssueDate     string `json:"CnicIssueDate"`
			CnicExpiryDate    string `json:"CnicExpiryDate"`
			DateOfBirth       string `json:"DateOfBirth"`
			MaritalStatus     string `json:"MaritalStatus"`
			ZakatStatus       string `json:"ZakatStatus"`
			Religion          string `json:"Religion"`
			Nationality       string `json:"Nationality"`
			DualNational      string `json:"DualNational"`
			ResidentialStatus string `json:"ResidentialStatus"`
			PhoneNo           string `json:"PhoneNo"`
			OfficeNo          string `json:"OfficeNo"`
			MobileNo          string `json:"MobileNo"`
			MobileNetwork     string `json:"MobileNetwork"`
			Ported            string `json:"Ported"`
			EmailAddress      string `json:"EmailAddress"`
			CurrentAddress    string `json:"CurrentAddress"`
			CurrentCity       string `json:"CurrentCity"`
			CurrentCountry    string `json:"CurrentCountry"`
			PermanentAddress  string `json:"PermanentAddress"`
			PermanentCity     string `json:"PermanentCity"`
			PermanentCountry  string `json:"PermanentCountry"`
			BankName          string `json:"BankName"`
			BankAccountNo     string `json:"BankAccountNo"`
			BranchName        string `json:"BranchName"`
			BranchCity        string `json:"BranchCity"`
			CashDividend      string `json:"CashDividend"`
			StockDividend     string `json:"StockDividend"`
			NextOfKin         struct {
				FullName      string `json:"FullName"`
				ContactNumber string `json:"ContactNumber"`
				Relationship  string `json:"Relationship"`
				Address       string `json:"Address"`
				Cnic          string `json:"Cnic"`
				AccountType   string `json:"AccountType"`
			} `json:"nextOfKin"`
			JointHolder struct {
				FullName     string `json:"FullName"`
				Relationship string `json:"Relationship"`
				Nic          string `json:"Nic"`
				IssueDate    string `json:"IssueDate"`
				ExpiryDate   string `json:"ExpiryDate"`
				Cnic         string `json:"Cnic"`
				AccountType  string `json:"AccountType"`
			} `json:"jointHolder"`
			SimOwner struct {
				SimOwnerType string      `json:"SimOwnerType"`
				SimOwnerCnic string      `json:"SimOwnerCnic"`
				Cnic         interface{} `json:"Cnic"`
				AccountType  interface{} `json:"AccountType"`
			} `json:"simOwner"`
			DaoID                string `json:"DaoID"`
			Cnic                 string `json:"Cnic"`
			AccountType          string `json:"AccountType"`
			Channel              string `json:"Channel"`
			AccountTransactionID string `json:"accountTransactionId"`
		} `json:"basicInfo"`
		HealthDec interface{} `json:"healthDec"`
		Kyc       struct {
			ResidentialStatus              string `json:"ResidentialStatus"`
			SourceOfIncome                 string `json:"SourceOfIncome"`
			SourceOfWealth                 string `json:"SourceOfWealth"`
			NameOfEmployer                 string `json:"NameOfEmployer"`
			Designation                    string `json:"Designation"`
			NatureOfBusiness               string `json:"NatureOfBusiness"`
			Education                      string `json:"Education"`
			Profession                     string `json:"Profession"`
			GeographiesDomestic            string `json:"GeographiesDomestic"`
			GeographiesIntl                string `json:"GeographiesIntl"`
			CounterPartyDomestic           string `json:"CounterPartyDomestic"`
			CounterPartyIntl               string `json:"CounterPartyIntl"`
			ModeOfTransaction              string `json:"ModeOfTransaction"`
			NumberOfTransaction            string `json:"NumberOfTransaction"`
			ExpectedTurnOverMonthlyAnnual  string `json:"ExpectedTurnOverMonthlyAnnual"`
			ExpectedTurnOver               string `json:"ExpectedTurnOver"`
			ExpectedInvestmentAmount       string `json:"ExpectedInvestmentAmount"`
			AnnualIncome                   string `json:"AnnualIncome"`
			ActionOnBehalfOfOther          int    `json:"actionOnBehalfOfOther"`
			RefusedYourAccount             int    `json:"refusedYourAccount"`
			SeniorPositionInGovtInstitute  int    `json:"seniorPositionInGovtInstitute"`
			SeniorPositionInPoliticalParty int    `json:"seniorPositionInPoliticalParty"`
			FinanciallyDependent           int    `json:"financiallyDependent"`
			HighValueGoldSilverDiamond     int    `json:"highValueGoldSilverDiamond"`
			IncomeIsHighRisk               int    `json:"incomeIsHighRisk"`
			HeadOfState                    int    `json:"headOfState"`
			SeniorMilitaryOfficer          int    `json:"seniorMilitaryOfficer"`
			HeadOfDeptOrIntlOrg            int    `json:"headOfDeptOrIntlOrg"`
			MemberOfBoard                  int    `json:"memberOfBoard"`
			MemberOfNationalSenate         int    `json:"memberOfNationalSenate"`
			PoliticalPartyOfficials        int    `json:"politicalPartyOfficials"`
			PepDeclaration                 string `json:"Pep_Declaration"`
			WhereDidYouHearAboutUs         string `json:"WhereDidYouHearAboutUs"`
			DaoID                          string `json:"DaoID"`
			Cnic                           string `json:"Cnic"`
			AccountType                    string `json:"AccountType"`
		} `json:"kyc"`
		Fatca struct {
			AccountTitle               string `json:"AccountTitle"`
			CountryOfTaxResidence      string `json:"CountryOfTaxResidence"`
			PobCity                    string `json:"PobCity"`
			PobState                   string `json:"PobState"`
			PobCountry                 string `json:"PobCountry"`
			AreYouUsCitizen            int    `json:"areYouUsCitizen"`
			AreYouUsResident           int    `json:"areYouUsResident"`
			HoldGreenCard              int    `json:"holdGreenCard"`
			WhereBornInUsa             int    `json:"whereBornInUsa"`
			AccountMaintainedInUsa     int    `json:"accountMaintainedInUsa"`
			PowerOfAttorney            int    `json:"powerOfAttorney"`
			UsResidenceMailiingAddress int    `json:"UsResidenceMailiingAddress"`
			UsTelephoneNumber          int    `json:"UsTelephoneNumber"`
			Cnic                       string `json:"Cnic"`
			AccountType                string `json:"AccountType"`
		} `json:"fatca"`
		Crs struct {
			CountryOfTaxResidence string `json:"CountryOfTaxResidence"`
			TinNo                 string `json:"TinNo"`
			Explaination          string `json:"Explaination"`
			City                  string `json:"City"`
			State                 string `json:"State"`
			Address               string `json:"Address"`
			Cnic                  string `json:"Cnic"`
			AccountType           string `json:"AccountType"`
		} `json:"crs"`
		RiskProfile struct {
			KeyValue            interface{} `json:"KeyValue"`
			AgeInYears          int         `json:"AgeInYears"`
			Age                 string      `json:"Age"`
			RiskReturn          string      `json:"RiskReturn"`
			MonthlySavings      string      `json:"MonthlySavings"`
			Occupation          string      `json:"Occupation"`
			InvestmentObjective string      `json:"InvestmentObjective"`
			KnowledgeLevel      string      `json:"KnowledgeLevel"`
			InvestmentHorizon   string      `json:"InvestmentHorizon"`
			IdealPortfolio      interface{} `json:"idealPortfolio"`
			IdealFund           interface{} `json:"idealFund"`
			IdealScore          interface{} `json:"idealScore"`
			Cnic                string      `json:"Cnic"`
			AccountType         string      `json:"AccountType"`
		} `json:"riskProfile"`
	} `json:"KeyValue"`
}
type CreateMeezanAccountResponseStruct []struct {
	ErrID  string `json:"ErrID"`
	ErrMsg string `json:"ErrMsg"`
}

type InvesetMeezanRequestStruct struct {
	Username        string `json:"Username"`
	Password        string `json:"Password"`
	TransactionType string `json:"TransactionType"`
	Key             string `json:"Key"`
	PortfolioID     string `json:"PortfolioID"`
	Branch          string `json:"Branch"`
	Account         string `json:"Account"`
	FundID          string `json:"FundId"`
	Amount          string `json:"Amount"`
	Unit            string `json:"Unit"`
}
type InvesetMeezanResponseStruct struct {
	OrderID  string `json:"OrderID"`
	DateTime string `json:"DateTime"`
}

type RedeemMeezanRequestStruct struct {
	Username        string `json:"Username"`
	Password        string `json:"Password"`
	TransactionType string `json:"TransactionType"`
	Key             string `json:"Key"`
	PortfolioID     string `json:"PortfolioID"`
	Branch          string `json:"Branch"`
	Account         string `json:"Account"`
	FundID          string `json:"FundId"`
	Amount          string `json:"Amount"`
	Unit            string `json:"Unit"`
}

type RedeemMeezanResponseStruct struct {
	OrderID  string `json:"OrderID"`
	DateTime string `json:"DateTime"`
}

type MeezanReversalRequestStruct struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Key      string `json:"Key"`
}

type MeezanReversalResponseStruct struct {
	Response        string `json:"response"`
	ResponseMessage string `json:"responseMessage"`
}

type MeezanDocUploadRequestStruct struct {
	Username             string `json:"Username"`
	Password             string `json:"Password"`
	Cnic                 string `json:"Cnic"`
	PrincipalCnicFront   string `json:"PrincipalCnicFront"`
	PrincipalCnicBack    string `json:"PrincipalCnicBack"`
	SourceOfIncome       string `json:"SourceOfIncome"`
	ZakatDeclaration     string `json:"ZakatDeclaration"`
	MobileBill           string `json:"MobileBill"`
	SignCard             string `json:"SignCard"`
	Affidavite           string `json:"Affidavite"`
	SimOwnerCnicFront    string `json:"SimOwnerCnicFront"`
	SimOwnerCnicBack     string `json:"SimOwnerCnicBack"`
	HealthDeclaration    string `json:"HealthDeclaration"`
	JointHolderCnicFront string `json:"JointHolderCnicFront"`
	JointHolderCnicBack  string `json:"JointHolderCnicBack"`
	JointHolderSignCard  string `json:"JointHolderSignCard"`
	LivePhoto            string `json:"LivePhoto"`
}

type MeezanDocUploadResponseStruct []struct {
	ErrID  string `json:"ErrID"`
	ErrMsg string `json:"ErrMsg"`
}

type MeezanConvertRequestStruct struct {
	Username        string `json:"Username"`
	Password        string `json:"Password"`
	TransactionType string `json:"TransactionType"`
	Key             string `json:"Key"`
	PortfolioID     string `json:"PortfolioID"`
	Branch          string `json:"Branch"`
	Account         string `json:"Account"`
	FundID          string `json:"FundId"`
	Amount          string `json:"Amount"`
	Unit            string `json:"Unit"`
	FundIDTo        string `json:"FundIdTo"`
}

type MeezanConvertResponseStruct struct {
	OrderID  string `json:"OrderID"`
	DateTime string `json:"DateTime"`
}
