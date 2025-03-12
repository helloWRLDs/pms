package scheduler

import (
	"context"
	"time"

	"pms.pkg/logger"
)

func Run(ctx context.Context, task *Task) {
	log := logger.Log.With("func", "scheduler.Run", "task_id", task.ID)
	log.Debug("scheduler.Run called")

	ticker := time.NewTicker(task.Interval)

	ctx, cancel := context.WithCancel(ctx)
	task.cancelFunc = cancel

	log.Debugw("Creating ticker", "interval", task.Interval)

	task.SetStatus(TASK_STATUS_RUNNING)

	go func() {
		defer func() {
			ticker.Stop()
		}()

		for {
			select {
			case <-ticker.C:
				switch task.Status() {
				case TASK_STATUS_PAUSED:
					continue
				case TASK_STATUS_DONE, TASK_STATUS_CANCELED:
					return
				default:
					if err := task.Func(ctx); err == nil {
						if task.MaxAttempts > 0 {
							task.SetStatus(TASK_STATUS_DONE)
							return
						}
					} else {
						log.Errorw("task failed", "error", err)
					}

					if task.MaxAttempts > 0 {
						task.AddAttempt()
						log.Warnw("task attempt failed", "attempt", task.Attempts(), "max_attempts", task.MaxAttempts)

						if task.Attempts() >= task.MaxAttempts {
							log.Warn("exceeded max attempts, marking as failed")
							task.SetStatus(TASK_STATUS_FAILED)
							return
						}
					}
				}
			case <-ctx.Done():
				log.Debug("task canceled via context")
				task.SetStatus(TASK_STATUS_CANCELED)
				return
			}
		}
	}()
}
