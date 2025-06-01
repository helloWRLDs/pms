package models

import (
	"pms.pkg/datastore/redis"
	"pms.pkg/transport/grpc/dto"
	ctxutils "pms.pkg/utils/ctx"
)

var _ redis.Cachable = &DocumentBody{}
var _ ctxutils.ContextKeyHolder = &DocumentBody{}

type DocumentBody struct {
	Document      *dto.Document `json:"document"`
	RequireUpdate bool
}

func (DocumentBody) GetDB() int {
	return 2
}
func (d DocumentBody) ContextKey() ctxutils.ContextKey {
	return ctxutils.ContextKey("document_body")
}
