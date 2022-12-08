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
}
