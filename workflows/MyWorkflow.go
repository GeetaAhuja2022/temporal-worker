package workflows

import (
	"fmt"
	"temporal-workflow/activities"

	"go.temporal.io/sdk/workflow"
)

// MyWorkflow is a Temporal workflow that uses the factory pattern to execute activities.
func MyWorkflow(ctx workflow.Context, activityType activities.ActivityType, input interface{}) (string, error) {
	factory := &activities.ActivityFactory{}

	// Create the desired activity
	activity, err := factory.CreateActivity(activityType)
	if err != nil {
		return "", err
	}

	// Execute the activity
	result, err := workflow.ExecuteActivity(ctx, activity.Execute, input).Get(ctx, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Workflow finished with result: %v", result), nil
}
