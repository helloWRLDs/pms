package utils

import (
	"pms.pkg/consts"
	"pms.pkg/transport/grpc/dto"
)

func CalculateTaskPoints(task *dto.Task) int32 {
	var points int32 = 0

	// Base points by task type
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

	// Get base points from task type
	if basePoints, exists := typePoints[task.Type]; exists {
		points += basePoints
	} else {
		points += 10 // Default points for unknown types
	}

	// Priority multiplier (1 = highest, 5 = lowest)
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

	// Completion time bonus/penalty (only for done tasks)
	if task.Status == string(consts.TASK_STATUS_DONE) {
		if task.DueDate != nil && task.UpdatedAt != nil {
			dueTime := task.DueDate.AsTime()
			completedTime := task.UpdatedAt.AsTime()

			// Early completion bonus
			if completedTime.Before(dueTime) {
				daysDiff := int(dueTime.Sub(completedTime).Hours() / 24)
				if daysDiff > 0 {
					// Bonus: 1 point per day early, max 10 points
					bonus := min(daysDiff, 10)
					points += int32(bonus)
				}
			} else if completedTime.After(dueTime) {
				// Late completion penalty
				daysDiff := int(completedTime.Sub(dueTime).Hours() / 24)
				if daysDiff > 0 {
					// Penalty: -1 point per day late, max -10 points
					penalty := min(daysDiff, 10)
					points -= int32(penalty)
				}
			}
		}

		// Quality bonus based on task age (faster completion = bonus)
		if task.CreatedAt != nil && task.UpdatedAt != nil {
			completionTime := task.UpdatedAt.AsTime().Sub(task.CreatedAt.AsTime())

			// Quick completion bonus (within 1 day)
			if completionTime.Hours() <= 24 {
				points += 5
			} else if completionTime.Hours() <= 72 { // Within 3 days
				points += 2
			}

			// Penalty for tasks taking too long (over 2 weeks)
			if completionTime.Hours() > 24*14 {
				points -= 3
			}
		}
	}

	// Ensure minimum points
	if points < 1 {
		points = 1
	}

	return points
}
