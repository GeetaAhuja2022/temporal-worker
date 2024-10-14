package activities

import (
	"context"
	"fmt"
)

type PaymentActivity struct{}

// PaymentActivity represents processing a payment.

// PaymentActivityFunction is the function to be registered in Temporal.
func PaymentActivityFunction(ctx context.Context, input map[string]interface{}) (string, error) {
	activity := &PaymentActivity{}
	result, err := activity.Execute(ctx, input)
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

// Execute performs the actual payment processing.
func (p *PaymentActivity) Execute(ctx context.Context, input interface{}) (interface{}, error) {
	paymentData, ok := input.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type for PaymentActivity")
	}
	fmt.Printf("Processing payment of $%.2f\n", paymentData["amount"])
	return "Payment Processed", nil
}
