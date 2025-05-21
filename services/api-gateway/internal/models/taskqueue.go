package models

import (
	"pms.pkg/datastore/redis"
	"pms.pkg/transport/grpc/dto"
)

var _ redis.Cachable = &TaskQueueElement{}

type TaskQueueElement struct {
	Value *dto.Task
}

func (tq TaskQueueElement) GetDB() int {
	return 1
}
