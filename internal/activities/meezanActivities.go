package activities

import (
	"Github/amc-apis-go/internal/data"
	"Github/amc-apis-go/internal/enums"
	"bytes"
	"context"
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

func (c *Client) CreateRequest(method, url string, RequestBody []byte, BearerToken string) (returnVal *http.Response, err error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(RequestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", BearerToken)

	res, err := c.hc.Do(req)
	if err != nil {
		return res, err
	}
	return res, nil
}

func GetToken() (ApiTokenReponse data.GetTokenStruct, errs error) {
	requestBody := data.RequestTokenStruct{Username: "admin", Password: "adminpass"}
	rawData, marshalErr := json.Marshal(requestBody)
	if marshalErr != nil {
		log.Fatal(marshalErr, nil)
	}

	req, err := http.NewRequest("POST", enums.GetTokenURL, bytes.NewBuffer(rawData))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := NewClient()
	res, err := client.hc.Do(req)

	defer res.Body.Close()

	ApiReponse, err := ioutil.ReadAll(res.Body)

	json.Unmarshal(ApiReponse, &ApiTokenReponse)

	log.Print(ApiTokenReponse)

	return ApiTokenReponse, err
}

func CreateAccount(ctx context.Context, UserData data.CreateAccountRequestStruct, BearerToken string) (returnVal data.CreateAccountResponseStruct, errs error) {

	reqData, err := json.Marshal(UserData)
	if err != nil {
		log.Fatal(err)
	}

	client := NewClient()

	apiRes, err := client.CreateRequest("POST", enums.CreateAccURL, reqData, BearerToken)

	if err != nil {
		log.Fatal(err)
	}

	defer apiRes.Body.Close()

	resData, err := ioutil.ReadAll(apiRes.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(resData, &returnVal)
	log.Print(returnVal)

	return returnVal, nil

}

func TestActivity(ctx context.Context, n1 int, n2 int) int {
	return n1 + n2
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
