package models

import "pms.pkg/transport/grpc/dto"

type DocumentBody struct {
	Document      *dto.Document `json:"document"`
	RequireUpdate bool
}

func (DocumentBody) GetDB() int {
	return 2
}
