package main

import (
	"fmt"
	hw05parallelexecution "github.com/uralmetal/hw_golang_otus/hw05_parallel_execution"
	"math/rand"
	"sync/atomic"
	"time"
)

func main() {
	tasksCount := 50
	tasks := make([]hw05parallelexecution.Task, 0, tasksCount)

	var runTasksCount int32
	var sumTime time.Duration

	for i := 0; i < tasksCount; i++ {
		taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
		sumTime += taskSleep

		tasks = append(tasks, func() error {
			time.Sleep(taskSleep)
			atomic.AddInt32(&runTasksCount, 1)
			fmt.Println("task", atomic.LoadInt32(&runTasksCount))
			return nil
		})
	}

	workersCount := 5
	maxErrorsCount := 1

	err := hw05parallelexecution.Run(tasks, workersCount, maxErrorsCount)
	fmt.Println("error:", err)
	fmt.Println("task count:", runTasksCount)
}
