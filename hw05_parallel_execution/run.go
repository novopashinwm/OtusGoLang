package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var errCnt int32
	taskCh := make(chan Task)
	wg := &sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range taskCh {
				if err := task; err != nil && m > 1 {
					atomic.AddInt32(&errCnt, 1)
				} else {
					task()
				}
			}
		}()
	}
	for _, task := range tasks {
		if m > 1 && atomic.LoadInt32(&errCnt) >= int32(m) {
			break
		}
		taskCh <- task
	}
	close(taskCh)
	wg.Wait()

	if errCnt >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
