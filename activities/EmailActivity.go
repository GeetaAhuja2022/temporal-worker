package activities

import (
	"context"
	"fmt"
)

type EmailActivity struct{}

// EmailActivityFunction is the function to be registered in Temporal.
func EmailActivityFunction(ctx context.Context, input map[string]string) (string, error) {
	activity := &EmailActivity{}
	result, err := activity.Execute(ctx, input)
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

// Execute performs the actual email sending.
func (e *EmailActivity) Execute(ctx context.Context, input interface{}) (interface{}, error) {
	emailData, ok := input.(map[string]string)
	if !ok {
		return nil, fmt.Errorf("invalid input type for EmailActivity")
	}
	fmt.Printf("Sending email to %s with subject: %s\n", emailData["to"], emailData["subject"])
	return "Email Sent", nil
}
