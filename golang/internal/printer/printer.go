package printer

import (
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/NodaSoft/tasks/internal/task"
)

type PrintFormat int

const (
	ListFormat   = PrintFormat(0)
	AmountFormat = PrintFormat(1)
)

type Printer struct {
	writer io.Writer
	format PrintFormat
}

func New(writer io.Writer, format PrintFormat) *Printer {
	return &Printer{writer: writer, format: format}
}

func (p *Printer) Print(period time.Duration, tasks chan task.Task) {
	successfulTasks := []task.Task{}
	failedTasks := []task.Task{}
	stopPrinter := false

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for {
			time.Sleep(period)

			_, err := p.writer.Write([]byte("Current state. Successful tasks:\n"))
			if err != nil {
				log.Println("Error writing header:", err)
			}

			if p.format == ListFormat {
				for _, task := range successfulTasks {
					_, err := p.writer.Write([]byte(task.String() + "\n"))
					if err != nil {
						log.Println("Error writing successful task:", err)
					}
				}
			} else if p.format == AmountFormat {
				_, err := p.writer.Write([]byte(fmt.Sprintln(len(successfulTasks))))
				if err != nil {
					log.Println("Error writing successful task:", err)
				}
			}

			_, err = p.writer.Write([]byte("Failed tasks:\n"))
			if err != nil {
				log.Println("Error writing failed header:", err)
			}

			if p.format == ListFormat {
				for _, task := range failedTasks {
					_, err := p.writer.Write([]byte(task.String() + "\n"))
					if err != nil {
						log.Println("Error writing failed task:", err)
					}
				}
			} else if p.format == AmountFormat {
				_, err := p.writer.Write([]byte(fmt.Sprintln(len(failedTasks))))
				if err != nil {
					log.Println("Error writing successful task:", err)
				}
			}

			if stopPrinter {
				break
			}
		}

		wg.Done()
	}()

	for currentTask := range tasks {
		if currentTask.IsSuccess() {
			successfulTasks = append(successfulTasks, currentTask)
		} else if currentTask.IsFail() {
			failedTasks = append(failedTasks, currentTask)
		}
	}

	stopPrinter = true
	wg.Wait()
}
