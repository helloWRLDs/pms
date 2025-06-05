package notifiermq

import (
	mqtp "pms.pkg/transport/mq"
)

var (
	Routes = []mqtp.QueueRoute{
		GreetMessage{}.RoutingKey(),
		TaskAssignmentMessage{}.RoutingKey(),
	}
)
