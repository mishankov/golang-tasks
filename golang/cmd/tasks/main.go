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
	printFormat              = printer.Amount
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	resultChan := make(chan task.Task)

	go func() {
		time.Sleep(tasksGenerationSeconds * time.Second)
		cancel()
	}()

	generator := generator.New()
	newTasksChan := generator.Generate(ctx)

	var wg sync.WaitGroup
	for range workersAmount {
		wg.Add(1)
		worker := worker.New()
		go worker.Work(&wg, newTasksChan, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	printer := printer.New(os.Stdout, printFormat)
	printer.Print(resultPrintPeriodSeconds*time.Second, resultChan)
}
