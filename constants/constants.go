package constants

// Serve configs
const (
	ProdEnvironment = "production"
	DevEnvironment  = "development"
	ApplicationPort = ":8080"
)

// Endpoints
const (
	TaskTrackerGroup    = "/tasktracker"
	TaskTrackerEndpoint = "/task"
)

// TaskIdParam params
const (
	TaskIdParam = ":taskId"
)

// Status
const (
	Started   = "STARTED"
	Pending   = "PENDING"
	Completed = "COMPLETED"
)

// Error Log

const (
	ErrWhileBinding    = "error while binding json: "
	ErrWhileValidating = "error while validating json: "
)
