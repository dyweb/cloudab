package v1

import (
	"github.com/google/uuid"
)

type Identity struct {
	ID     *uuid.UUID `json:"id,omitempty"`
	UserID *string    `json:"user_id,omitempty"`
	// TODO(gaocegege): Support Session.
}

type Experiment struct {
	ID   *uuid.UUID `json:"id,omitempty"`
	Name string     `json:"name,omitempty"`

	// TODO(gaocegege): Support layers.
	Versions []EnvironmentVersion `json:"versions,omitempty"`
	Metrics  []Metric             `json:"metrics,omitempty"`
}

type EnvironmentVersion struct {
	ID   *uuid.UUID `json:"id,omitempty"`
	Name string     `json:"name,omitempty"`

	Traffic  int32     `json:"traffic,omitempty"`
	Features []Feature `json:"features,omitempty"`
}

type Feature struct {
	ID *uuid.UUID `json:"id,omitempty"`

	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`

	// TODO(gaocegege): Support number, boolean.
}

type Metric struct {
	ID        *uuid.UUID `json:"id,omitempty"`
	Name      string     `json:"id,omitempty"`
	EventName string     `json:"event_name,omitempty"`
}

// Message describes a message entry.
type Message struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
