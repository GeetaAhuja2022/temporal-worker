package activities

import (
	"context"
	"fmt"
)

// EmailActivity represents sending an email.
type EmailActivity struct{}

// EmailActivityFunction is the registered function in Temporal.
func EmailActivityFunction(ctx context.Context, input map[string]string) (ActivityResult, error) {
	activity := &EmailActivity{}
	result, err := activity.Execute(ctx, input)
	if err != nil {
		return ActivityResult{Status: "error", Message: err.Error()}, err
	}
	return result.(ActivityResult), nil
}

// Execute sends an email and returns a common activity result.
func (e *EmailActivity) Execute(ctx context.Context, input interface{}) (interface{}, error) {
	emailData, ok := input.(map[string]string)
	if !ok {
		return ActivityResult{}, fmt.Errorf("invalid input type for EmailActivity")
	}
	// Simulate sending an email
	fmt.Printf("Sending email to %s with subject: %s\n", emailData["To"], emailData["Subject"])

	// Return a standardized response
	return ActivityResult{
		Status:  "success",
		Message: "Email Sent",
		Data:    nil, // You can return any additional data here
	}, nil
}
