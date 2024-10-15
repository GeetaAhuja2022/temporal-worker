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

	emailInput := map[string]string{
		"to":      "user@example.com",
		"subject": "Welcome to Temporal",
	}
	paymentInput := map[string]string{
		"Amount": "1000",
	}
	input := map[string]map[string]string{
		"emailData":   emailInput,
		"paymentData": paymentInput,
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflows.MyWorkflow, input)
	if err != nil {
		log.Fatalln("Unable to start workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
