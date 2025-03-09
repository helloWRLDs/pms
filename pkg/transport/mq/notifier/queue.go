package notifiermq

import mqtp "pms.pkg/transport/mq"

const (
	Queue mqtp.QueueName = "notifier"
)

type MetaData struct {
	ToEmail string `json:"to_email"`
}
