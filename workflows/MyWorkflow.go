package workflows

import (
	"fmt"
	"temporal-workflow/activities"

	"go.temporal.io/sdk/workflow"
)

func MyWorkflow(ctx workflow.Context, input interface{}) (string, error) {
	var result string
	var err error

	// Dynamically execute the correct activity based on type
	factory = activities.ActivityFactory()
	var emailActivityInstance = activities.ActivityFactory.CreateActivity(0)
	var paymentActivityInstance = activities.ActivityFactory.CreateActivity(1)

	err = workflow.ExecuteActivity(ctx, activities.EmailActivityFunction, input).Get(ctx, &result)
	err = workflow.ExecuteActivity(ctx, activities.PaymentActivityFunction, input).Get(ctx, &result)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Workflow finished with result: %v", result), nil
}
