package task

import (
	"fmt"
	"time"
)

type Result string

const (
	Success = Result("Success")
	Fail    = Result("Fail")
)

type Task struct {
	Id         uint64
	CreateTime time.Time
	FinishTime time.Time
	Result     Result
	Error      error
}

func New() Task {
	return Task{
		Id:         uint64(time.Now().UnixNano()),
		CreateTime: time.Now(),
	}
}

func (t Task) IsSuccess() bool {
	return t.Result == Success
}

func (t Task) IsFail() bool {
	return t.Result == Fail
}

func (t Task) String() string {
	if t.IsFail() {
		return fmt.Sprintf("%v - %v (%v)", t.Id, t.Result, t.Error)
	}
	return fmt.Sprintf("%v - %v", t.Id, t.Result)
}
