package iniservice

import (
	"taskTracker/core/relationaldatabase"
	"taskTracker/internal/adapters/egress/dao"
	"taskTracker/internal/ports/tasktraceregress"
)

type Adapter struct {
	tasktraceregress.DbOps
}

func NewAdapter(dbWrapper *relationaldatabase.DbWrapper, funcs ...tasktraceregress.DbOpsFunc) *Adapter {
	ops := GetDefaults(dbWrapper)
	for _, fn := range funcs {
		fn(&ops)
	}
	return &Adapter{DbOps: ops}
}

func GetDefaults(wrapper *relationaldatabase.DbWrapper) tasktraceregress.DbOps {
	return tasktraceregress.DbOps{
		TaskTracker: dao.NewTrackerAdapter(wrapper),
		User:        dao.NewUserAdapter(wrapper),
	}
}

func (adapter Adapter) GetTaskTracker() tasktraceregress.TaskTracker {
	return adapter.TaskTracker
}

func (adapter Adapter) GetUser() tasktraceregress.User {
	return adapter.User
}
