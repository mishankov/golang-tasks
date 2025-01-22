package printer_test

import (
	"bytes"
	"errors"
	"testing"
	"time"

	"github.com/NodaSoft/tasks/internal/printer"
	"github.com/NodaSoft/tasks/internal/task"
)

func TestPrinter(t *testing.T) {
	t.Parallel()

	buffer := &bytes.Buffer{}
	p := printer.New(buffer, printer.ListFormat)
	tasksChan := make(chan task.Task)

	go func() {
		tasksChan <- task.Task{Id: 1, Result: task.Success}
		tasksChan <- task.Task{Id: 2, Result: task.Fail, Error: errors.New("some error")}
		close(tasksChan)
	}()

	p.Print(1*time.Second, tasksChan)

	want := `Current state. Successful tasks:
1 - Success
Failed tasks:
2 - Fail (some error)
`

	if buffer.String() != want {
		t.Error("buffer incorrect, got:", buffer.String())
	}
}
