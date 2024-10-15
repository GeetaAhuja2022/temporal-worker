package activities

// ActivityResult is a common return type for all activities.
type ActivityResult struct {
	Status  string      // Example: "success" or "error"
	Message string      // A human-readable message
	Data    interface{} // Any additional data the activity might return
}
