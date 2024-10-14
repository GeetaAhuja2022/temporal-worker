package main

import (
	"log"
	"temporal-workflow/activities"
	"temporal-workflow/workflows"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Create a Temporal client
	c, err := client.NewLazyClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	// Create a worker for the task queue
	w := worker.New(c, "my-task-queue", worker.Options{})

	// Register the workflow
	w.RegisterWorkflow(workflows.MyWorkflow)

	// Register the activities
	w.RegisterActivity(activities.EmailActivityFunction)
	w.RegisterActivity(activities.PaymentActivityFunction)

	// Start the worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
