package scheduler

type TaskStatus int

const (
	TASK_STATUS_RUNNING TaskStatus = iota
	TASK_STATUS_DONE
	TASK_STATUS_CANCELED
	TASK_STATUS_PAUSED
	TASK_STATUS_FAILED
)

func (ts TaskStatus) String() string {
	switch ts {
	case 0:
		return "RUNNING"
	case 1:
		return "DONE"
	case 2:
		return "CANCELED"
	case 3:
		return "PAUSED"
	case 4:
		return "FAILED"
	}
	return ""
}
