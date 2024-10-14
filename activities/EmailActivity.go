package activities

import (
	"context"
	"fmt"
)

// EmailActivity represents sending an email.
type EmailActivity struct{}

// Execute is the implementation for sending an email.
func (e *EmailActivity) Execute(ctx context.Context, input interface{}) (interface{}, error) {
	emailData, ok := input.(map[string]string)
	if !ok {
		return nil, fmt.Errorf("invalid input type for EmailActivity")
	}
	// Simulate sending an email.
	fmt.Printf("Sending email to %s with subject: %s\n", emailData["to"], emailData["subject"])
	return "Email Sent", nil
}
