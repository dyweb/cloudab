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
	Versions []Version `json:"versions,omitempty" bson:"versions,omitempty"`
	Metrics  []Metric  `json:"metrics,omitempty" bson:"metrics,omitempty"`
}

type Version struct {
	ID   *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string              `json:"name,omitempty" bson:"name,omitempty"`

	Traffic  int32     `json:"traffic,omitempty" bson:"traffic,omitempty"`
	Features []Feature `json:"features,omitempty" bson:"features,omitempty"`
}

type Feature struct {
	ID *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`

	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Value string `json:"value,omitempty" bson:"value,omitempty"`

	// TODO(gaocegege): Support number, boolean.
}

type Metric struct {
	ID        *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string              `json:"name,omitempty" bson:"name,omitempty"`
	EventName string              `json:"event_name,omitempty" bson:"event_name,omitempty"`
}

// Message describes a message entry.
type Message struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
