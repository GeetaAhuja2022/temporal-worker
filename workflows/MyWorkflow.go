package workflows

import (
	"fmt"
	"temporal-workflow/activities"

	"go.temporal.io/sdk/workflow"
)

// MyWorkflow is a Temporal workflow that uses a factory to execute activities with a common return type.
func MyWorkflow(ctx workflow.Context, input map[string]interface{}) (string, error) {
	var result activities.ActivityResult
	var err error

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
