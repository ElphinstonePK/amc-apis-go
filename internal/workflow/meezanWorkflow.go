package myWorkflows

import (
	"Github/amc-apis-go/internal/activities"
	"Github/amc-apis-go/internal/data"
	"Github/amc-apis-go/internal/database"
	"log"
	"os"
	"time"

	"go.temporal.io/sdk/workflow"
)

func MeezanBalanceInquiryWorkflow(ctx workflow.Context, input data.BalanceInquiryRequestStruct) (balanceInquiryResponse data.BalanceInquiryResponseStruct, errs error) {

	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:              "meezan",
		ScheduleToCloseTimeout: 20 * time.Second,
		//ScheduleToStartTimeout: 3 * time.Second,
		StartToCloseTimeout: 20 * time.Second,
		HeartbeatTimeout:    0 * time.Second,
		WaitForCancellation: false,
	})

	tokenErr := workflow.ExecuteActivity(ctx, activities.GetMeezanAccountBalance, input).Get(ctx, &balanceInquiryResponse)

	if tokenErr != nil {
		log.Print("Failure on executing GetToken activity")
		log.Fatal(tokenErr, nil)
	}

	return balanceInquiryResponse, tokenErr
}

func MeezanInvestWorkflow(ctx workflow.Context, input data.InvesetMeezanRequestStruct) (investmentResponse data.InvesetMeezanResponseStruct, errs error) {

	/*
		1. Ping AMC API
		2. Add to the DB the CNIC + trasnactionStatus + fundID + the current amount of units in that fund

		3. Add a loop that keeps pinging the BalanceInquiry API to check against the store amount
		of units in the DBto check if update has been made to the user's balance

		4. Update the DB, transactionStatus to complete
		5. Send the confirmation info B2B/B2C & Update units in B2B/B2C
	*/

	errs = database.ConnectToDB(
		os.Getenv("DATABASE_URL"),
		"meezan",
		"investments",
	)

	if errs != nil {
		log.Fatal(errs)
	}

	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:              "meezan",
		ScheduleToCloseTimeout: 20 * time.Second,
		//ScheduleToStartTimeout: 3 * time.Second,
		StartToCloseTimeout: 20 * time.Second,
		HeartbeatTimeout:    0 * time.Second,
		WaitForCancellation: false,
	})

	tokenErr := workflow.ExecuteActivity(ctx, activities.InevestMeezan, input).Get(ctx, &investmentResponse)

	if tokenErr != nil {
		log.Print("Failure on executing GetToken activity")
		log.Fatal(tokenErr, nil)
	}

	errs = database.CreateInvestmentRecord(
		input.PortfolioID,
		input.FundID,
	)

	return investmentResponse, tokenErr
}
