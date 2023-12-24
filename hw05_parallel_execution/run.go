package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
var ErrInvalidParameters = errors.New("errors invalid parameters")

type Task func() error

func worker(tasks []Task, m *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	for _, task := range tasks {
		err := task()
		if err != nil {
			atomic.AddInt64(m, -1)
		}
		if atomic.LoadInt64(m) <= 0 {
			break
		}
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	errorCount := (int64)(m)
	workersTasks := make(map[int][]Task)

	if m <= 0 || n <= 0 {
		return ErrErrorsLimitExceeded
	}
	for i, task := range tasks {
		workersTasks[i%n] = append(workersTasks[i%n], task)
	}
	for i := range workersTasks {
		if len(workersTasks) > 0 {
			//wg.Add(1)
			go worker(workersTasks[i], &errorCount, &wg)
		}
	}
	wg.Wait()
	if errorCount <= 0 {
		return ErrErrorsLimitExceeded
	}
	return nil
}
