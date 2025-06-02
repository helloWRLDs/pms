package utils

import (
	"pms.pkg/consts"
	"pms.pkg/transport/grpc/dto"
)

func CalculateTaskPoints(task *dto.Task) int32 {
	var points int32 = 0

	typePoints := map[string]int32{
		string(consts.TaskTypeFeature):       25, // Features are most valuable
		string(consts.TaskTypeStory):         20, // User stories are high value
		string(consts.TaskTypeBug):           15, // Bug fixes are important
		string(consts.TaskTypeRefactor):      12, // Code improvements
		string(consts.TaskTypeTest):          10, // Testing tasks
		string(consts.TaskTypeChore):         8,  // Maintenance tasks
		string(consts.TaskTypeSubTask):       5,  // Sub-tasks are smaller
		string(consts.TaskTypeDocumentation): 6,  // Documentation work
	}

	if basePoints, exists := typePoints[task.Type]; exists {
		points += basePoints
	} else {
		points += 10
	}

	priorityMultipliers := map[int32]float32{
		1: 1.5,  // Highest priority: +50%
		2: 1.25, // High priority: +25%
		3: 1.0,  // Medium priority: base
		4: 0.8,  // Low priority: -20%
		5: 0.6,  // Lowest priority: -40%
	}

	if multiplier, exists := priorityMultipliers[task.Priority]; exists {
		points = int32(float32(points) * multiplier)
	}

	if task.Status == string(consts.TASK_STATUS_DONE) {
		if task.DueDate != nil && task.UpdatedAt != nil {
			dueTime := task.DueDate.AsTime()
			completedTime := task.UpdatedAt.AsTime()

			if completedTime.Before(dueTime) {
				daysDiff := int(dueTime.Sub(completedTime).Hours() / 24)
				if daysDiff > 0 {

					bonus := min(daysDiff, 10)
					points += int32(bonus)
				}
			} else if completedTime.After(dueTime) {

				daysDiff := int(completedTime.Sub(dueTime).Hours() / 24)
				if daysDiff > 0 {

					penalty := min(daysDiff, 10)
					points -= int32(penalty)
				}
			}
		}

		if task.CreatedAt != nil && task.UpdatedAt != nil {
			completionTime := task.UpdatedAt.AsTime().Sub(task.CreatedAt.AsTime())

			if completionTime.Hours() <= 24 {
				points += 5
			} else if completionTime.Hours() <= 72 {
				points += 2
			}

			if completionTime.Hours() > 24*14 {
				points -= 3
			}
		}
	}

	if points < 1 {
		points = 1
	}

	return points
}
