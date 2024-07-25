package main

import (
	"cloud.google.com/go/bigquery"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	credentialsFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	dataset := os.Getenv("BIGQUERY_DATASET")

	fmt.Printf("files: %s, %s, %s\n", credentialsFile, projectID, dataset) // TODO
	ctx := context.Background()
	if err := control(ctx, credentialsFile, projectID, dataset); err != nil {
		fmt.Fprintf(os.Stderr, "fail in controling: %v", err)
		os.Exit(1)
	}
}

func control(ctx context.Context, credentialsFile, projectID, dataset string) error {
	cli, err := bigquery.NewClient(ctx, projectID, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %v", err)
	}
	schemaMap, err := loadSchemas(ctx, cli, dataset)
	if err != nil {
		return fmt.Errorf("loadSchemas: %v", err)
	}
	for table, schema := range schemaMap {
		for _, field := range schema {
			fmt.Println(table, field.Name, field.Type)
		}
	}
	return nil
}

func loadSchemas(ctx context.Context, cli *bigquery.Client, dataset string) (map[string]bigquery.Schema, error) {
	schemaMap := make(map[string]bigquery.Schema)
	t := cli.Dataset(dataset).Tables(ctx)
	for {
		table, err := t.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("iterator.Next: %v", err)
		}
		meta, err := table.Metadata(ctx)
		if err != nil {
			return nil, fmt.Errorf("metadata: %v", err)
		}
		schemaMap[table.TableID] = meta.Schema
	}
	return schemaMap, nil

}
