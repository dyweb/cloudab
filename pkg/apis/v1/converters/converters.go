package converters

import (
	"context"

	"github.com/caicloud/nirvana/log"
	"github.com/caicloud/nirvana/operators/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ObjectID is the converter from HEX string to ObjectID.
var ObjectID = converter.For(
	func(ctx context.Context, field string, value string) (primitive.ObjectID, error) {
		logger := log.DefaultLogger()

		id, err := primitive.ObjectIDFromHex(value)
		if err != nil {
			logger.V(log.LevelDebug).Infof(
				"Failed to convert HEX string %s to ObjectID: %v", value, err)
		}
		return id, err
	})
