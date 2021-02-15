package metrics

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
