package activities

import (
	"context"
	"fmt"
)

// PaymentActivity represents processing a payment.
type PaymentActivity struct{}

// Execute is the implementation for processing a payment.
func (p *PaymentActivity) Execute(ctx context.Context, input interface{}) (interface{}, error) {
	paymentData, ok := input.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type for PaymentActivity")
	}
	amount := paymentData["amount"]
	fmt.Printf("Processing payment of $%.2f\n", amount)
	return "Payment Processed", nil
}
