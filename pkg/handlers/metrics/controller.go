package metrics

import (
	"context"
	"fmt"
	"time"

	"github.com/caicloud/nirvana/log"
	v1 "github.com/dyweb/cloudab/pkg/apis/v1"
	"github.com/dyweb/cloudab/pkg/store"
	"github.com/dyweb/cloudab/pkg/traffic"
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

// ReportMetrics reports metrics.
func (c Controller) ReportMetrics(ctx context.Context,
	id primitive.ObjectID, userID string,
	metricName string, value int64) (*v1.Experiment, error) {
	logger := log.DefaultLogger()

	cctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	singleResult := store.DB.Collection(c.collection).FindOne(cctx, bson.D{
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
	logger.V(log.LevelDebug).Infof("Routing in %d version(s) for the given experiment %s",
		len(versions), experiment.ID.Hex())
	r := traffic.NewRouter()
	versionID, err := r.Route(userID, versions)
	if err != nil {
		return nil, err
	}

	// Check if the versionID works as expected.
	if len(versions) <= versionID {
		return nil, fmt.Errorf("Failed to get the version")
	} else if versions[versionID].ID == nil {
		return nil, fmt.Errorf("Failed to get the version at %d: %v",
			versionID, versions[versionID])
	}

	if versions[versionID].Metrics == nil {
		versions[versionID].Metrics = make(map[string]v1.Metric)
	}
	versions[versionID].Metrics[metricName] = v1.Metric{
		Name:  metricName,
		Value: value,
	}

	cctx, cancel = context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	// TODO(gaocegege): Use updateOnce instead.
	_, err = store.DB.Collection(c.collection).ReplaceOne(cctx,
		bson.D{
			{
				Key:   "_id",
				Value: id,
			},
		}, experiment)
	if err != nil {
		return nil, err
	}

	return &experiment, nil
}
