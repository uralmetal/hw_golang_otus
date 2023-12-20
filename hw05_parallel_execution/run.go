package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func worker(task Task, output chan error, wg *sync.WaitGroup) {
	wg.Add(1)
	output <- task()
	wg.Done()
}

func workerPool(input chan Task, output chan error, wg *sync.WaitGroup) {
	for task := range input {
		go worker(task, output, wg)
	}
	wg.Wait()
}

func launcher(tasks []Task, input chan Task) {
	for task := range tasks {
		input <- task
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	var firstAppendCount int
	input := make(chan Task, n)
	output := make(chan error, 1)

	wg.Add(1)
	go launcher(tasks, input)
	go workerPool(input, output, &wg)
	if len(tasks) > n {
		firstAppendCount = n
	} else {
		firstAppendCount = len(tasks)
	}
	for i := 0; i < firstAppendCount-1; i++ {
		input <- tasks[i]
	}
	wg.Done()
	for i := firstAppendCount; i < len(tasks); i++ {
		taskError := <-output
		if taskError != nil {
			m--
		}
		if m <= 0 {
			return ErrErrorsLimitExceeded
		}
		input <- tasks[i]
	}

	return nil
}
