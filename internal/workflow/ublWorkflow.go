package myWorkflows

import (
	"Github/amc-apis-go/internal/activities"
	"Github/amc-apis-go/internal/data"
	"log"
	"time"

	"go.temporal.io/sdk/workflow"
)

func UblCreateAccWorkflow(ctx workflow.Context, input data.CreateAccountRequestStruct) (data.CreateAccountResponseStruct, error) {
	var tokenResponse data.GetTokenStruct
	var createAccResponse data.CreateAccountResponseStruct

	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:              "ubl",
		ScheduleToCloseTimeout: 20 * time.Second,
		//ScheduleToStartTimeout: 3 * time.Second,
		StartToCloseTimeout: 20 * time.Second,
		HeartbeatTimeout:    0 * time.Second,
		WaitForCancellation: false,
	})

	tokenErr := workflow.ExecuteActivity(ctx, activities.GetUBLToken).Get(ctx, &tokenResponse)

	if tokenErr != nil {
		log.Print("Failure on executing GetToken activity")
		log.Fatal(tokenErr, nil)
	}

	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:              "ubl",
		ScheduleToCloseTimeout: 20 * time.Second,
		//ScheduleToStartTimeout: 3 * time.Second,
		StartToCloseTimeout: 20 * time.Second,
		HeartbeatTimeout:    0 * time.Second,
		WaitForCancellation: false,
	})

	createAccError := workflow.ExecuteActivity(ctx, activities.CreateUBLAccount, input, tokenResponse.AccessToken).Get(ctx, &createAccResponse)

	if createAccError != nil {
		log.Print("Failure on executing CreateAccount activity")
		log.Fatal(createAccError, nil)
	}

	return createAccResponse, nil
}

func TestWorkflow(ctx workflow.Context, n1 int, n2 int) (testReponse int, errs error) {
	var testResponse int

	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:              "test",
		ScheduleToCloseTimeout: 20 * time.Second,
		//ScheduleToStartTimeout: 3 * time.Second,
		StartToCloseTimeout: 20 * time.Second,
		HeartbeatTimeout:    0 * time.Second,
		WaitForCancellation: false,
	})

	testErr := workflow.ExecuteActivity(ctx, activities.TestActivity, n1, n2).Get(ctx, &testResponse)

	if testErr != nil {
		log.Print("Failure on executing Test activity")
		log.Fatal(testErr, nil)
	}

	return testReponse, errs
}
