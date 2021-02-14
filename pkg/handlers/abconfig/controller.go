package abconfig

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/caicloud/nirvana/log"
	v1 "github.com/dyweb/cloudab/pkg/apis/v1"
	"github.com/dyweb/cloudab/pkg/store"
	"github.com/dyweb/cloudab/pkg/traffic"
)

type Controller struct {
	experimentCollection string
}

// New creates a new Controller.
func New() Controller {
	return Controller{
		experimentCollection: "experiments",
	}
}

// GetABConfig return an experiment by id.
func (c Controller) GetABConfig(ctx context.Context,
	id primitive.ObjectID, userID string) (*v1.ABConfig, error) {
	logger := log.DefaultLogger()

	cctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	singleResult := store.DB.Collection(c.experimentCollection).FindOne(cctx, bson.D{
		{
			Key:   "_id",
			Value: id,
		},
	})
	if singleResult == nil {
		return nil, fmt.Errorf("Failed to find the experiment with id %s", id.Hex())
	}
	var experiment v1.Experiment
	if err := singleResult.Decode(&experiment); err != nil {
		return nil, err
	}

	versions := experiment.Versions
	logger.V(log.LevelDebug).Infof(
		"Routing in %d version(s) for the given experiment %s", len(versions), experiment.ID.Hex())
	r := traffic.NewRouter()
	versionID, err := r.Route(userID, versions)
	if err != nil {
		return nil, err
	}

	// Check if the versionID works as expected.
	if len(versions) <= versionID {
		return nil, fmt.Errorf("Failed to get the version")
	} else if versions[versionID].ID == nil {
		return nil, fmt.Errorf("Failed to get the version at %d: %v", versionID, versions[versionID])
	}
	return &v1.ABConfig{
		ExperimentName: experiment.Name,
		ExperimentID:   experiment.ID,
		Versions: []string{
			versions[versionID].ID.Hex(),
		},
	}, nil
}
