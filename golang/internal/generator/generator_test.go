package generator_test

import (
	"context"
	"testing"
	"time"

	"github.com/NodaSoft/tasks/internal/generator"
)

func TestGenerate(t *testing.T) {
	t.Parallel()

	gen := generator.New()
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	newTasksChan := gen.Generate(ctx)

	generatedTasksAmount := 0
	for currentTask := range newTasksChan {
		generatedTasksAmount++
		if currentTask.Id == 0 {
			cancel()
			t.Errorf("got invalid task: %v", currentTask)
		}
	}

	if generatedTasksAmount == 0 {
		t.Error("no tasks were generated")
	}
}
