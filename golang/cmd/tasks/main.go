package main

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/NodaSoft/tasks/internal/generator"
	"github.com/NodaSoft/tasks/internal/printer"
	"github.com/NodaSoft/tasks/internal/task"
	"github.com/NodaSoft/tasks/internal/worker"
)

const (
	workersAmount            = 4
	tasksGenerationSeconds   = 10
	resultPrintPeriodSeconds = 3
	printFormat              = printer.ListFormat
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), tasksGenerationSeconds*time.Second)
	defer cancel()

	procesedTasks := make(chan task.Task)

	generator := generator.New()
	generatedTasks := generator.Generate(ctx)

	var wg sync.WaitGroup
	for range workersAmount {
		wg.Add(1)
		worker := worker.New()
		go worker.Work(&wg, generatedTasks, procesedTasks)
	}

	go func() {
		wg.Wait()
		close(procesedTasks)
	}()

	printer := printer.New(os.Stdout, printFormat)
	printer.Print(resultPrintPeriodSeconds*time.Second, procesedTasks)
}
