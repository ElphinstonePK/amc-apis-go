package main

import (
	"Github/amc-apis-go/internal/data"
	myWorkflows "Github/amc-apis-go/internal/workflow"
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{HostPort: "15.184.170.196:7233"})
	if err != nil {
		log.Fatal(err, nil)
	}
	defer c.Close()
	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: "meezan",
	}

	var investmentObj data.InvesetMeezanRequestStruct
	investmentObj.Username = "elphinstone"
	investmentObj.Password = "DLRr7G@ksn%xa$!5"
	investmentObj.TransactionType = "Investment"
	investmentObj.Key = "MBLTransactionID"
	investmentObj.PortfolioID = "186506-1"
	investmentObj.Branch = "0103"
	investmentObj.Account = "020282501"
	investmentObj.FundID = "800102-100~100120"
	investmentObj.Amount = "50000"
	investmentObj.Unit = "0"

	_, err = c.ExecuteWorkflow(context.Background(), workflowOptions, myWorkflows.MeezanInvestWorkflow, investmentObj)
	if err != nil {
		log.Fatal(err, nil)
	}
}
