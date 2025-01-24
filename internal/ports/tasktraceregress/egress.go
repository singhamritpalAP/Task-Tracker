package tasktraceregress

type DbPort interface {
	GetTaskTracker() TaskTracker
}

type TaskTracker interface {
	Create() error
	Update() error
	Read() error
	Delete() error
}
