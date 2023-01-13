package main

import (
	activities "Github/amc-apis-go/internal/activities"
	myWorkflows "Github/amc-apis-go/internal/workflow"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func register(w worker.Registry) {
	w.RegisterWorkflowWithOptions(
		myWorkflows.TestWorkflow,
		workflow.RegisterOptions{Name: "Test Workflow"},
	)

	w.RegisterActivityWithOptions(
		activities.TestActivity,
		activity.RegisterOptions{Name: "Test Activity"},
	)

	w.RegisterWorkflowWithOptions(
		myWorkflows.MeezanBalanceInquiryWorkflow,
		workflow.RegisterOptions{Name: "Meezan Account Inquiry Workflow"},
	)

	w.RegisterActivityWithOptions(
		activities.GetMeezanAccountBalance,
		activity.RegisterOptions{Name: "Meezan Account Inquiry Activity"},
	)

	w.RegisterWorkflowWithOptions(
		myWorkflows.MeezanInvestWorkflow,
		workflow.RegisterOptions{Name: "Meezan Invest Workflow"},
	)

	w.RegisterActivityWithOptions(
		activities.InevestMeezan,
		activity.RegisterOptions{Name: "Meezan Invest Activity"},
	)
}
