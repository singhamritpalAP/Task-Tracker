package tasktrackerigress

type TaskTrackerAPIPort interface {
	GetTaskAPIPort() TaskAPIPort
}

type TaskAPIPort interface {
	CreateTask() error
}
