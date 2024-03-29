package v1

import "go.mongodb.org/mongo-driver/bson/primitive"

type Identity struct {
	ID     *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID *string             `json:"user_id,omitempty" bson:"user_id,omitempty"`
	// TODO(gaocegege): Support Session.
}

type Experiment struct {
	ID   *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string              `json:"name,omitempty" bson:"name,omitempty"`

	// TODO(gaocegege): Support layers.
	// TODO(gaocegege): Should we keep versions in another collection?
	Versions []Version `json:"versions,omitempty" bson:"versions,omitempty"`
}

type Version struct {
	ID   *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string              `json:"name,omitempty" bson:"name,omitempty"`

	Traffic  int32     `json:"traffic,omitempty" bson:"traffic,omitempty"`
	Features []Feature `json:"features,omitempty" bson:"features,omitempty"`

	// TODO(gaocegege): Should we keep metrics in another collection?
	Metrics map[string]Metric `json:"metrics,omitempty" bson:"metrics,omitempty"`
}

type Feature struct {
	ID *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`

	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Value string `json:"value,omitempty" bson:"value,omitempty"`

	// TODO(gaocegege): Support number, boolean.
}

type ABConfig struct {
	Features       []Feature           `json:"features,omitempty" bson:"features,omitempty"`
	ExperimentName string              `json:"experiment_name,omitempty" bson:"experiment_name,omitempty"`
	ExperimentID   *primitive.ObjectID `json:"experiment_id,omitempty" bson:"experiment_id,omitempty"`
	Versions       []string            `json:"versions,omitempty" bson:"versions,omitempty"`
}

type Metric struct {
	ID        *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string              `json:"name,omitempty" bson:"name,omitempty"`
	EventName string              `json:"event_name,omitempty" bson:"event_name,omitempty"`
	Value     int64               `json:"value,omitempty"`
}
