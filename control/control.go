package control

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/koooyooo/bq-schema/output"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type Options struct {
	ExcludeTablePatterns string
}

func Control(ctx context.Context, credentialsFile, projectID, dataset string, opts *Options) error {
	cli, err := bigquery.NewClient(ctx, projectID, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %v", err)
	}
	schemaMap, err := loadSchemas(ctx, cli, dataset)
	if err != nil {
		return fmt.Errorf("loadSchemas: %v", err)
	}
	excludePatterns := strings.Split(opts.ExcludeTablePatterns, "::")
	for i, pattern := range excludePatterns {
		excludePatterns[i] = strings.TrimSpace(pattern)
	}
	for tableName, _ := range schemaMap {
		for _, excludePattern := range excludePatterns {
			matched, err := regexp.MatchString(excludePattern, tableName)
			if err != nil {
				return fmt.Errorf("regexp.MatchString: %v", err)
			}
			if matched {
				delete(schemaMap, tableName)
			}
		}
	}

	outputDir := "target"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("os.Mkdir: %v", err)
	}
	formatter := output.FindFormatter(output.FormatterOptionPlantUML)
	files, err := formatter(ctx, schemaMap)
	if err != nil {
		return fmt.Errorf("f: %v", err)
	}
	for _, file := range files {
		if err := os.WriteFile(outputDir+"/"+file.Name, file.Content, 0644); err != nil {
			return fmt.Errorf("os.WriteFile: %v", err)
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
