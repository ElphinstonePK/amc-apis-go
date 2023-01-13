package data

type AccountOpeningDBStruct struct {
	Cnic   string `bson:"_id"`
	Status string `bson:"status"`
}

type InvestmentDBStruct struct {
	Cnic   string  `bson:"_id"`
	Status string  `bson:"status"`
	FundId string  `bson:"fund_id"`
	Units  float64 `bson:"units"`
}
