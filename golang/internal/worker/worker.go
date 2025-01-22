package worker

import (
	"sync"
	"time"

	"github.com/NodaSoft/tasks/internal/task"
)

type Worker struct{}

func New() *Worker {
	return &Worker{}
}

func (w *Worker) Work(wg *sync.WaitGroup, inputChan chan task.Task, resultChan chan task.Task) {
	for currentTask := range inputChan {
		time.Sleep(time.Millisecond * 150)
		if currentTask.Error != nil {
			currentTask.Result = task.Fail
		} else {
			currentTask.Result = task.Success
		}

		currentTask.FinishTime = time.Now()

		resultChan <- currentTask
	}

	wg.Done()
}
