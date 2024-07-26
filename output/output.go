package output

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/bigquery"
)

type FormatterOption int

const (
	FormatterOptionPlantUML FormatterOption = iota
)

type File struct {
	Name    string
	Content []byte
}

type Formatter func(ctx context.Context, schemas map[string]bigquery.Schema) ([]*File, error)

func FindFormatter(opt FormatterOption) Formatter {
	switch opt {
	case FormatterOptionPlantUML:
		return FormatterPlantUML
	default:
		return FormatterPlantUML
	}
}

func FormatterPlantUML(ctx context.Context, schemas map[string]bigquery.Schema) ([]*File, error) {
	var b bytes.Buffer
	b.WriteString("@startuml\n")
	for title, schema := range schemas {
		b.WriteString(fmt.Sprintf("entity %s {\n", title))
		for _, field := range schema {
			b.WriteString(fmt.Sprintf("\t+ %s: %s\n", field.Name, strings.ToLower(string(field.Type))))
		}
		b.WriteString(fmt.Sprintln("}"))
	}
	b.WriteString("@enduml\n\n")
	return []*File{{
		Name:    "schemas.puml",
		Content: b.Bytes(),
	}}, nil
}
