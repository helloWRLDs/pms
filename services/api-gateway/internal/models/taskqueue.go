package models

import (
	"pms.pkg/datastore/redis"
	"pms.pkg/transport/grpc/dto"
	ctxutils "pms.pkg/utils/ctx"
)

var _ redis.Cachable = &TaskQueueElement{}
var _ ctxutils.ContextKeyHolder = &TaskQueueElement{}

type TaskQueueElement struct {
	Value *dto.Task `json:"value"`
}

func (tq TaskQueueElement) GetDB() int {
	return 1

}
func (tq TaskQueueElement) ContextKey() ctxutils.ContextKey {
	return ctxutils.ContextKey("task_queue")
}
