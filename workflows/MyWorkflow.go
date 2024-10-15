package workflows

import (
	"fmt"
	"time"

	"github.com/GeetaAhuja2022/temporal-worker/activities"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// MyWorkflow is a Temporal workflow that uses a factory to execute activities with a common return type.
func MyWorkflow(ctx workflow.Context, input map[string]interface{}) (string, error) {

	retryPolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    time.Minute,
		MaximumAttempts:    3,
	}

	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failures by default, this is just an example.
		RetryPolicy: retryPolicy,
	}
	var result activities.ActivityResult
	var err error
	ctx = workflow.WithActivityOptions(ctx, options)

	// Dynamically execute the correct activity based on type

	err = workflow.ExecuteActivity(ctx, activities.EmailActivityFunction, input["emailData"]).Get(ctx, &result)

	if err != nil {
		return "", err
	}

	err = workflow.ExecuteActivity(ctx, activities.PaymentActivityFunction, input["paymentData"]).Get(ctx, &result)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Workflow finished with status: %s, message: %s", result.Status, result.Message), nil
}
