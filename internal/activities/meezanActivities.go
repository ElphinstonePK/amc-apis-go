package activities

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	""
)

type Client struct {
	hc *http.Client
}

func    NewClient() *Client {
	c := &http.Client{}
	return &Client{hc: c}
}

func (c *Client) CreateRequest(method, url string, RequestBody []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(RequestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := c.hc.Do(req)
	if err != nil {
		return res, err
	}
	return res, nil
}

func CreatePaymentSafepay(ctx context.Context, input data.FinalPaymentRequestStruct) (SafepayCreatePaymentResponse data.CreatePaymentResponseStruct, errs error) {

	safepayReq := data.CreatePaymentRequestStructSafepay{
		Amount:      input.Amount,
		Client:      "sec_d64f8626-d820-465b-857e-684afea3c56f",
		Currency:    "PKR",
		Environment: "sandbox",
	}

	rawData, err := json.Marshal(safepayReq)
	if err != nil {
		log.Fatal(err, nil)
	}

	c := NewClient()

	res, err := c.CreateRequest("POST", enums.CreatePaymentURLSafepay, rawData)
	if err != nil {
		log.Fatal(err, nil)
	}

	defer res.Body.Close()

	SafepayApiResponse, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err, nil)
	}

	json.Unmarshal(SafepayApiResponse, &SafepayCreatePaymentResponse)

	log.Print(SafepayCreatePaymentResponse)

	return SafepayCreatePaymentResponse, nil
}
