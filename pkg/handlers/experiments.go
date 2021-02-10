package handlers

import (
	"context"

	v1 "github.com/abtest-io/cloudab-server/pkg/apis/v1"
)

// ListExperiments returns all experiments.
func ListExperiments(ctx context.Context, count int) ([]v1.Message, error) {
	return nil, nil
}

func CreateExperiment(
	ctx context.Context, exp *v1.Experiment) (*v1.Experiment, error) {
	return exp, nil
}

// GetMessage return a message by id.
func GetMessage(ctx context.Context, id int) (*v1.Message, error) {
	return &v1.Message{
		ID:      id,
		Title:   "This is an example",
		Content: "Example content",
	}, nil
}
