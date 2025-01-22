package worker_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/NodaSoft/tasks/internal/generator"
	"github.com/NodaSoft/tasks/internal/task"
	"github.com/NodaSoft/tasks/internal/worker"
)

func TestWorker(t *testing.T) {
	t.Parallel()

	w := worker.New()
	ctx, cancel := context.WithCancel(context.Background())

	gen := generator.New()
	newTasksChan := gen.Generate(ctx)

	resultCh := make(chan task.Task)

	go func() {
		time.Sleep(1 * time.Second)
		cancel()
		close(resultCh)
	}()

	var wg sync.WaitGroup
	go w.Work(&wg, newTasksChan, resultCh)

	processedTasksAmount := 0
	for processedTask := range resultCh {
		processedTasksAmount++
		if processedTask.Result != task.Success && processedTask.Result != task.Fail {
			cancel()
			t.Errorf("task has invalid result: %q", processedTask.Result)
		}
	}

	if processedTasksAmount == 0 {
		t.Error("no task was processed")
	}
}
