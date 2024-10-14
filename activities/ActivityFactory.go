package activities

import "fmt"

// ActivityType defines the type of activity
type ActivityType string

const (
	EmailActivityType   ActivityType = "email"
	PaymentActivityType ActivityType = "payment"
)

// ActivityFactory is responsible for creating activities.
type ActivityFactory struct{}

// CreateActivity creates an activity instance based on the activity type.
func (af *ActivityFactory) CreateActivity(activityType ActivityType) (Activity, error) {
	switch activityType {
	case EmailActivityType:
		return &EmailActivity{}, nil
	case PaymentActivityType:
		return &PaymentActivity{}, nil
	default:
		return nil, fmt.Errorf("unknown activity type: %s", activityType)
	}
}
