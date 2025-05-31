package consts

type TaskType string

const (
	TaskTypeBug           TaskType = "bug"
	TaskTypeStory         TaskType = "story"
	TaskTypeSubTask       TaskType = "sub_task"
	TaskTypeFeature       TaskType = "feature"
	TaskTypeChore         TaskType = "chore"
	TaskTypeRefactor      TaskType = "refactor"
	TaskTypeTest          TaskType = "test"
	TaskTypeDocumentation TaskType = "documentation"
)

func (t TaskType) String() string {
	return string(t)
}
