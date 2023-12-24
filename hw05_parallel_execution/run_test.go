package hw05parallelexecution

import (
	"errors"
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestRun(t *testing.T) {
	defer goleak.VerifyNone(t)

	t.Run("if were errors in first M tasks, than finished not more N+M tasks", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32

		for i := 0; i < tasksCount; i++ {
			err := fmt.Errorf("error from task %d", i)
			tasks = append(tasks, func() error {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
				atomic.AddInt32(&runTasksCount, 1)
				return err
			})
		}

		workersCount := 10
		maxErrorsCount := 23
		err := Run(tasks, workersCount, maxErrorsCount)

		require.Truef(t, errors.Is(err, ErrErrorsLimitExceeded), "actual err - %v", err)
		require.LessOrEqual(t, runTasksCount, int32(workersCount+maxErrorsCount), "extra tasks were started")
	})

	t.Run("tasks less workers", func(t *testing.T) {
		tasksCount := 5
		tasks := make([]Task, 0, tasksCount)

		runTasksCount := atomic.Int32{}
		var sumTime time.Duration

		for i := 0; i < tasksCount; i++ {
			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
			sumTime += taskSleep

			tasks = append(tasks, func() error {
				runTasksCount.Add(1)
				time.Sleep(taskSleep)
				return nil
			})
		}

		workersCount := 50
		maxErrorsCount := 1

		err := Run(tasks, workersCount, maxErrorsCount)
		require.NoError(t, err)
		require.Equal(t, runTasksCount.Load(), int32(tasksCount), "not all tasks were completed")
	})

	t.Run("handle corrected invalid parameters", func(t *testing.T) {
		tasks := make([]Task, 0, 1)

		err := Run(tasks, 1, 0)
		require.Error(t, err, ErrInvalidParameters)

		err = Run(tasks, 0, 1)
		require.Error(t, err, ErrInvalidParameters)

		err = Run(tasks, 1, 1)
		require.NoError(t, err)
	})

	t.Run("tasks without errors", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32
		var sumTime time.Duration

		for i := 0; i < tasksCount; i++ {
			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
			sumTime += taskSleep

			tasks = append(tasks, func() error {
				atomic.AddInt32(&runTasksCount, 1)
				time.Sleep(taskSleep)
				return nil
			})
		}

		workersCount := 5
		maxErrorsCount := 1

		require.Eventually(t, func() bool {
			err := Run(tasks, workersCount, maxErrorsCount)
			if err != nil {
				return false
			}
			return true
		}, sumTime, time.Millisecond*100)

		require.Equal(t, runTasksCount, int32(tasksCount), "not all tasks were completed")
	})
}
