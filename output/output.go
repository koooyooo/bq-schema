package output

import (
	"bytes"
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
)

type File struct {
	Name    string
	Content []byte
}

type Formatter func(ctx context.Context, title string, schema bigquery.Schema) (*File, error)

func FormatterPlantUML(ctx context.Context, title string, schema bigquery.Schema) (*File, error) {
	var b bytes.Buffer
	b.WriteString("@startuml\n")
	b.WriteString(fmt.Sprintf("entity %s {\n", title))
	for _, field := range schema {
		b.WriteString(fmt.Sprintf("\t+ %s: %s\n", field.Name, field.Type))
	}
	b.WriteString(fmt.Sprintln("}"))
	b.WriteString("@enduml\n")
	return &File{
		Name:    title + ".puml",
		Content: b.Bytes(),
	}, nil
}
