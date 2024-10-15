package factory

import (
	"fmt"

	"github.com/GeetaAhuja2022/temporal-worker/activities"
)

// ActivityType defines the type of activity
type activityIType int

const (
	emailActivity activityIType = iota
	paymentActivity
)

var activityTypes = map[activityIType]string{
	emailActivity:   "EmailActivity",
	paymentActivity: "PaymentActivity",
}

// ActivityFactory is responsible for creating activities.
type ActivityFactory struct{}

// CreateActivity creates an activity instance based on the activity type.
func (af *ActivityFactory) CreateActivity(aType activityIType) (any, error) {
	switch aType {
	case emailActivity:
		return activities.EmailActivityFunction, nil
	case paymentActivity:
		return activities.PaymentActivityFunction, nil
	default:
		return nil, fmt.Errorf("unknown activity type: %s", aType)
	}
}
