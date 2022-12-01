package myWorkflows

import (
	"Github/amc-apis-go/internal/data"
	"log"
	"payment-apis/internal/activities"
	"payment-apis/internal/data"
	"time"

	"go.temporal.io/sdk/workflow"
)

func MeezanWorkflow(ctx workflow.Context) (data.GetTokenStruct, error) {
	var response data.GetTokenStruct

	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:              "meezan",
		ScheduleToCloseTimeout: 20 * time.Second,
		//ScheduleToStartTimeout: 3 * time.Second,
		StartToCloseTimeout: 20 * time.Second,
		HeartbeatTimeout:    0 * time.Second,
		WaitForCancellation: false,
	})

	err := workflow.ExecuteActivity(ctx, activities.GetToken).Get(ctx, &response)

	if err != nil {
		log.Print("Failure on executing activity")
		log.Fatal(err, nil)
	}

	return response, nil
}
