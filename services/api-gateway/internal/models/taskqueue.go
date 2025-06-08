package models

import (
	"pms.pkg/datastore/redis"
	"pms.pkg/transport/grpc/dto"
)

var _ redis.Cachable = &TaskQueueElement{}

type TaskQueueElement struct {
	Tasks         map[string]*dto.Task `json:"tasks"`
	TasksToUpdate []*dto.Task          `json:"tasks_to_update"`
}

func (tq TaskQueueElement) GetDB() int {
	return 1

}
