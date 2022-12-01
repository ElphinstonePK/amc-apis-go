package activities

import (
	"Github/amc-apis-go/internal/data"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	hc *http.Client
}

func NewClient() *Client {
	c := &http.Client{}
	return &Client{hc: c}
}

func (c *Client) GetToken() (ApiTokenReponse data.GetTokenStruct, errs error) {
	requestBody := data.RequestTokenStruct{Username: "admin", Password: "adminpass"}
	rawData, marshalErr := json.Marshal(requestBody)
	if marshalErr != nil {
		log.Fatal(marshalErr, nil)
	}
	req, err := http.NewRequest("POST", "/Token", bytes.NewBuffer(rawData))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := c.hc.Do(req)

	defer res.Body.Close()

	ApiReponse, err := ioutil.ReadAll(res.Body)

	json.Unmarshal(ApiReponse, &ApiTokenReponse)

	log.Print(ApiTokenReponse)

	return ApiTokenReponse, err
}

// func CreatePaymentSafepay(ctx context.Context, input data.FinalPaymentRequestStruct) (SafepayCreatePaymentResponse data.CreatePaymentResponseStruct, errs error) {

// 	safepayReq := data.CreatePaymentRequestStructSafepay{
// 		Amount:      input.Amount,
// 		Client:      "sec_d64f8626-d820-465b-857e-684afea3c56f",
// 		Currency:    "PKR",
// 		Environment: "sandbox",
// 	}

// 	rawData, err := json.Marshal(safepayReq)
// 	if err != nil {
// 		log.Fatal(err, nil)
// 	}

// 	c := NewClient()

// 	res, err := c.CreateRequest("POST", enums.CreatePaymentURLSafepay, rawData)
// 	if err != nil {
// 		log.Fatal(err, nil)
// 	}

// 	defer res.Body.Close()

// 	SafepayApiResponse, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		log.Fatal(err, nil)
// 	}

// 	json.Unmarshal(SafepayApiResponse, &SafepayCreatePaymentResponse)

// 	log.Print(SafepayCreatePaymentResponse)

// 	return SafepayCreatePaymentResponse, nil
// }
