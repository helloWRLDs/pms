package scheduler

import (
	"context"
	"fmt"
	"testing"
	"time"

	"pms.pkg/errs"
	"pms.pkg/tools/httpclient"
)

// readTaskStatus safely reads the status of the Task under its mutex.
// func readTaskStatus(t *Task) TaskStatus {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()
// 	return t.Status()
// }

func Test_Cancelation(t *testing.T) {
	task := Task{
		Func: func(ctx context.Context) error {
			return errs.ErrNotFound{
				Object: "user",
				Value:  "1",
				Field:  "id",
			}
		},
		ID:          "cancelation-task-1",
		MaxAttempts: 1000,
		Interval:    5 * time.Second,
	}
	Run(context.Background(), &task)
	time.Sleep(20 * time.Second)
	task.Cancel()
	time.Sleep(5 * time.Second)
	t.Log(task.Status())
}

func TestPingGoogleUsingScheduler(t *testing.T) {
	pingGoogle := func(ctx context.Context) error {
		res, err := httpclient.New().
			Method("GET").
			URL("https://google.com").
			Do()
		if err != nil {
			return err
		}

		if res.Status != 200 {
			return fmt.Errorf("pingGoogle: unexpected status code: %d", res.Status)
		}
		return nil
	}
	if err := pingGoogle(context.Background()); err != nil {
		t.Fatal(err)
	}
	t.Log("pinged google manually")

	task := &Task{
		ID:          "ping-google-task",
		Func:        pingGoogle,
		Interval:    1 * time.Second,
		MaxAttempts: 3,
	}

	Run(context.Background(), task)

	deadline := time.Now().Add(10 * time.Second)
LOOP:
	for time.Now().Before(deadline) {
		switch task.Status() {
		case TASK_STATUS_DONE:
			t.Log("Success!")
			break LOOP
		case TASK_STATUS_FAILED:
			t.Fatal("Task failed.")
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}

	if task.Status() != TASK_STATUS_DONE {
		t.Fatalf("Task did not complete by the deadline. Final status: %v", task.Status())
	}
}
