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

func (c *Client) CreateRequestMeezan(method, url string, RequestBody []byte) (returnVal *http.Response, err error) {
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

func GetMeezanAccountBalance(ctx context.Context, requestData data.BalanceInquiryRequestStruct) (finalResponseObj data.BalanceInquiryResponseStruct, errs error) {

	// log.Println("\n\n INPUT: ", requestData)

	rawData, marshalErr := json.Marshal(requestData)
	if marshalErr != nil {
		log.Fatal(marshalErr, nil)
	}

	myClient := NewClient()
	clientReponse, createReqError := myClient.CreateRequestMeezan("POST", enums.BalanceInquiry, rawData)
	if createReqError != nil {
		log.Fatal(createReqError, nil)
	}

	// log.Println("\n\n CLIENT RESPONSE: ", clientReponse)

	MeezanApiRepons, ioErr := ioutil.ReadAll(clientReponse.Body)

	// log.Println("\n\n IO RESPONSE: ", MeezanApiRepons)

	if ioErr != nil {
		log.Fatal(ioErr, nil)
	}

	json.Unmarshal(MeezanApiRepons, &finalResponseObj)

	log.Println("\n\n FINAL RESPONSE: ", finalResponseObj)

	defer clientReponse.Body.Close()

	return finalResponseObj, nil
}

func CreateMeezanAccount(ctx context.Context, requestData data.CreateMeezanAccountRequestStruct) (finalResponseObj data.CreateMeezanAccountResponseStruct, errs error) {
	rawData, marshalErr := json.Marshal(requestData)
	if marshalErr != nil {
		log.Fatal(marshalErr, nil)
	}

	myClient := NewClient()
	clientReponse, createReqError := myClient.CreateRequestMeezan("POST", enums.AccountOpening, rawData)
	if createReqError != nil {
		log.Fatal(createReqError, nil)
	}

	defer clientReponse.Body.Close()

	MeezanApiRepons, ioErr := ioutil.ReadAll(clientReponse.Body)

	if ioErr != nil {
		log.Fatal(ioErr, nil)
	}

	json.Unmarshal(MeezanApiRepons, &finalResponseObj)

	log.Print(finalResponseObj)

	return finalResponseObj, nil
}

func InevestMeezan(ctx context.Context, requestData data.InvesetMeezanRequestStruct) (finalResponseObj data.InvesetMeezanResponseStruct, errs error) {
	rawData, marshalErr := json.Marshal(requestData)
	if marshalErr != nil {
		log.Fatal(marshalErr, nil)
	}

	myClient := NewClient()
	clientReponse, createReqError := myClient.CreateRequestMeezan("POST", enums.Invest, rawData)
	if createReqError != nil {
		log.Fatal(createReqError, nil)
	}

	defer clientReponse.Body.Close()

	MeezanApiRepons, ioErr := ioutil.ReadAll(clientReponse.Body)

	if ioErr != nil {
		log.Fatal(ioErr, nil)
	}

	json.Unmarshal(MeezanApiRepons, &finalResponseObj)

	log.Println("\n\n INVESTMENT RESPONSE OBJ: ", finalResponseObj)

	return finalResponseObj, nil
}

func RedeemMeezan(ctx context.Context, requestData data.RedeemMeezanRequestStruct) (finalResponseObj data.RedeemMeezanResponseStruct, errs error) {
	rawData, marshalErr := json.Marshal(requestData)
	if marshalErr != nil {
		log.Fatal(marshalErr, nil)
	}

	myClient := NewClient()
	clientReponse, createReqError := myClient.CreateRequestMeezan("POST", enums.Redeem, rawData)
	if createReqError != nil {
		log.Fatal(createReqError, nil)
	}

	defer clientReponse.Body.Close()

	MeezanApiRepons, ioErr := ioutil.ReadAll(clientReponse.Body)

	if ioErr != nil {
		log.Fatal(ioErr, nil)
	}

	json.Unmarshal(MeezanApiRepons, &finalResponseObj)

	log.Println("\n\n REDEMPTION RESPONSE OBJ: ", finalResponseObj)

	return finalResponseObj, nil
}

func ReversalMeezan(ctx context.Context, requestData data.MeezanReversalRequestStruct) (finalResponseObj data.MeezanReversalResponseStruct, errs error) {
	rawData, marshalErr := json.Marshal(requestData)
	if marshalErr != nil {
		log.Fatal(marshalErr, nil)
	}

	myClient := NewClient()
	clientReponse, createReqError := myClient.CreateRequestMeezan("POST", enums.Reversal, rawData)
	if createReqError != nil {
		log.Fatal(createReqError, nil)
	}

	defer clientReponse.Body.Close()

	MeezanApiRepons, ioErr := ioutil.ReadAll(clientReponse.Body)

	if ioErr != nil {
		log.Fatal(ioErr, nil)
	}

	json.Unmarshal(MeezanApiRepons, &finalResponseObj)

	log.Println("\n\n REVERSAL RESPONSE OBJ: ", finalResponseObj)

	return finalResponseObj, nil
}

func DocUploadMeezan(ctx context.Context, requestData data.MeezanDocUploadRequestStruct) (finalResponseObj data.MeezanDocUploadResponseStruct, errs error) {
	rawData, marshalErr := json.Marshal(requestData)
	if marshalErr != nil {
		log.Fatal(marshalErr, nil)
	}

	myClient := NewClient()
	clientReponse, createReqError := myClient.CreateRequestMeezan("POST", enums.DocumentUpload, rawData)
	if createReqError != nil {
		log.Fatal(createReqError, nil)
	}

	defer clientReponse.Body.Close()

	MeezanApiRepons, ioErr := ioutil.ReadAll(clientReponse.Body)

	if ioErr != nil {
		log.Fatal(ioErr, nil)
	}

	json.Unmarshal(MeezanApiRepons, &finalResponseObj)

	log.Println("\n\n DOC UPLOAD RESPONSE OBJ: ", finalResponseObj)

	return finalResponseObj, nil
}

func ConvertFundsMeezan(ctx context.Context, requestData data.MeezanConvertRequestStruct) (finalResponseObj data.MeezanConvertResponseStruct, errs error) {
	rawData, marshalErr := json.Marshal(requestData)
	if marshalErr != nil {
		log.Fatal(marshalErr, nil)
	}

	myClient := NewClient()
	clientReponse, createReqError := myClient.CreateRequestMeezan("POST", enums.Conversion, rawData)
	if createReqError != nil {
		log.Fatal(createReqError, nil)
	}

	defer clientReponse.Body.Close()

	MeezanApiRepons, ioErr := ioutil.ReadAll(clientReponse.Body)

	if ioErr != nil {
		log.Fatal(ioErr, nil)
	}

	json.Unmarshal(MeezanApiRepons, &finalResponseObj)

	log.Println("\n\n CONVERT FUNDS RESPONSE OBJ: ", finalResponseObj)

	return finalResponseObj, nil
}
