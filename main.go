package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/koooyooo/bq-schema/control"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	credentialsFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	dataset := os.Getenv("BIGQUERY_DATASET")

	ctx := context.Background()
	if err := control.Control(ctx, credentialsFile, projectID, dataset); err != nil {
		fmt.Fprintf(os.Stderr, "fail in controling: %v", err)
		os.Exit(1)
	}
}
