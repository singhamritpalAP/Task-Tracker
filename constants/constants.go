package constants

// Serve configs
const (
	ProdEnvironment = "production"
	DevEnvironment  = "development"
	ApplicationPort = ":8080"
	JwtKey          = "taskTrackerSecureKey"
)

// Endpoints
const (
	TaskTrackerGroup    = "/tasktracker"
	TaskTrackerEndpoint = "/task"
	UserEndpoint        = "/user"
	LoginEndpoint       = "/login"
)

// TaskIdParam params
const (
	TaskIdParam = "taskId"
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

// Db where clauses
const (
	WhereId = "task_id = ?"
)

// General
const (
	TaskId = "task_id"
)
