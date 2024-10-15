package activities

import (
	"context"
	"fmt"
)

// PaymentActivity processes payments.
type PaymentActivity struct{}

// PaymentActivityFunction is the registered function in Temporal.
func PaymentActivityFunction(ctx context.Context, input map[string]interface{}) (ActivityResult, error) {
	activity := &PaymentActivity{}
	result, err := activity.Execute(ctx, input)
	if err != nil {
		return ActivityResult{Status: "error", Message: err.Error()}, err
	}
	return result.(ActivityResult), nil
}

// Execute processes a payment and returns a standardized result.
func (p *PaymentActivity) Execute(ctx context.Context, input interface{}) (interface{}, error) {
	paymentData, ok := input.(map[string]interface{})
	if !ok {
		return ActivityResult{}, fmt.Errorf("invalid input type for PaymentActivity")
	}

	// Simulate payment processing
	fmt.Printf("Processing payment of $%.2f\n", paymentData["Amount"])

	// Return a standardized response
	return ActivityResult{
		Status:  "success",
		Message: "Payment Processed",
		Data:    paymentData["amount"], // You can return additional data here
	}, nil
}
