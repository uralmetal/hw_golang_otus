package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrInvalidParameters   = errors.New("errors invalid parameters")
)

type Task func() error

func worker(tasks chan Task, m *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		err := task()
		if err != nil {
			atomic.AddInt64(m, -1)
		}
		if atomic.LoadInt64(m) <= 0 {
			break
		}
	}
}

func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	errorCount := (int64)(m)
	workersTasks := make(chan Task, len(tasks))

	if m <= 0 || n <= 0 {
		return ErrErrorsLimitExceeded
	}
	for _, task := range tasks {
		workersTasks <- task
	}
	close(workersTasks)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(workersTasks, &errorCount, &wg)
	}
	wg.Wait()
	if errorCount <= 0 {
		return ErrErrorsLimitExceeded
	}
	return nil
}
