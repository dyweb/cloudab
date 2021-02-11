package handlers

import (
	"context"

	v1 "github.com/dyweb/cloudab/pkg/apis/v1"
)

type name struct {
}

// ListExperiments returns all experiments.
func ListExperiments(ctx context.Context, count int) ([]v1.Message, error) {
	return nil, nil
}

// GetMessage return a message by id.
func GetMessage(ctx context.Context, id int) (*v1.Message, error) {
	return &v1.Message{
		ID:      id,
		Title:   "This is an example",
		Content: "Example content",
	}, nil
}
