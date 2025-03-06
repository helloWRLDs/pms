package scheduler

import (
	"context"
	"sync"
	"time"
)

type Task struct {
	ID          string
	MaxAttempts int
	Func        func(ctx context.Context) error
	Interval    time.Duration

	cancelFunc context.CancelFunc
	attempts   int
	status     TaskStatus

	mu sync.Mutex
}

func (t *Task) Attempts() int {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.attempts
}

func (t *Task) AddAttempt() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.attempts++
}

func (t *Task) Status() TaskStatus {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.status
}

func (t *Task) SetStatus(status TaskStatus) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.status = status
}

func (t *Task) Cancel() {
	t.cancelFunc()
}
