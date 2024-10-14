package main

import (
	"log"
	"temporal-workflow/workflows"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Create a Temporal client
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	// Create a worker for a workflow task queue
	w := worker.New(c, "my-task-queue", worker.Options{})

	// Register the workflow
	w.RegisterWorkflow(workflows.MyWorkflow)

	// Start the worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
