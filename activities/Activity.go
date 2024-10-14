package activities

import "context"

// Activity is an interface for all Temporal activities.
type SampleActivity interface {
	Execute(ctx context.Context, input interface{}) (interface{}, error)
}
