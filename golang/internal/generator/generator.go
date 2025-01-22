package generator

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/NodaSoft/tasks/internal/task"
)

type Generator struct{}

func New() *Generator {
	return &Generator{}
}

func (g *Generator) Generate(ctx context.Context) chan task.Task {
	ch := make(chan task.Task)
	totalTasks := 0

	go func() {
		for {
			newTask := task.New()
			if time.Now().Nanosecond()%2 > 0 { // TODO: always false on windows
				newTask.Error = errors.New("some error occured")
			}

			select {
			case <-ctx.Done():
				close(ch)
				log.Println("Generation stopped. Total tasks:", totalTasks)
				return
			case ch <- newTask:
				totalTasks++
			}
		}
	}()

	return ch
}
