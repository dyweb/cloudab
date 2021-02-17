package experiments

import (
	"context"
	"fmt"
	"time"

	"github.com/caicloud/nirvana/log"
	v1 "github.com/dyweb/cloudab/pkg/apis/v1"
	"github.com/dyweb/cloudab/pkg/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Controller defines the Experiment controller which is used to handle user requests.
type Controller struct {
	collection string
}

// New creates a new Controller.
func New() Controller {
	return Controller{
		collection: "experiments",
	}
}

// CreateExperiment creates a new experiment.
func (c Controller) CreateExperiment(
	ctx context.Context, exp *v1.Experiment) (*v1.Experiment, error) {
	cctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	logger := log.DefaultLogger()
	logger.V(log.LevelDebug).Infof(
		"Creating the experiment %s in collection %s", exp.Name, c.collection)

	// Set objectID for versions.
	for i := range exp.Versions {
		versionID := primitive.NewObjectID()
		exp.Versions[i].ID = &versionID
	}

	result, err := store.DB.Collection(c.collection).InsertOne(cctx, exp)
	if err != nil {
		return nil, err
	}
	// Set id for the retured object.
	id := result.InsertedID.(primitive.ObjectID)
	exp.ID = &id
	return exp, nil
}

// ListExperiments returns all experiments.
func (c Controller) ListExperiments(ctx context.Context, count int) ([]v1.Experiment, error) {
	cctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	var exp v1.Experiment
	exps, err := store.DB.Collection(c.collection).Find(cctx, exp)
	if err != nil {
		return nil, err
	}

	results := &[]v1.Experiment{}
	if err := exps.All(cctx, results); err != nil {
		return nil, err
	}
	return *results, nil
}

// GetExperiment return an experiment by id.
func (c Controller) GetExperiment(ctx context.Context, id primitive.ObjectID) (*v1.Experiment, error) {
	cctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	exp := store.DB.Collection(c.collection).FindOne(cctx, bson.D{
		{
			Key:   "_id",
			Value: id,
		},
	})
	if exp == nil {
		return nil, fmt.Errorf("Failed to find the experiment with id %s", id.Hex())
	}
	var result v1.Experiment
	if err := exp.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
