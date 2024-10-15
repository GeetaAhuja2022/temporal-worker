package main

import (
	"context"
	"log"
	"temporal-workflow/workflows"

	"go.temporal.io/sdk/client"
)

func main() {
	// Create a Temporal client
	c, err := client.NewLazyClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	// Start a workflow execution
	workflowOptions := client.StartWorkflowOptions{
		ID:        "factory-workflow",
		TaskQueue: "my-task-queue",
	}

	input := map[string]interface{}{
		"emailData": {
			"to":      "user@example.com",
			"subject": "Welcome to Temporal",
		},
		"paymentData": {
			"amount": "1000",
		},
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflows.MyWorkflow, input)
	if err != nil {
		log.Fatalln("Unable to start workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
