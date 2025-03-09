package notifiermq

import (
	mqtp "pms.pkg/transport/mq"
)

var (
	_ mqtp.Queueable = &GreetMessage{}
)

type GreetMessage struct {
	mqtp.MetaData
	Name string `json:"name"`
}

func (c GreetMessage) RoutingKey() mqtp.QueueRoute {
	return mqtp.QueueRoute("greet")
}
