package notifiermq

import (
	mqtp "pms.pkg/transport/mq"
)

var (
	_ mqtp.Queueable = &GreetMessage{}
	_ mqtp.Queueable = &TaskAssignmentMessage{}
)

type GreetMessage struct {
	MetaData
	Name string `json:"name"`
}

func (c GreetMessage) RoutingKey() mqtp.QueueRoute {
	return mqtp.QueueRoute("greet")
}

type TaskAssignmentMessage struct {
	MetaData
	AssigneeName string `json:"assignee_name"`
	TaskName     string `json:"task_name"`
	TaskId       string `json:"task_id"`
	ProjectName  string `json:"project_name"`
}

func (c TaskAssignmentMessage) RoutingKey() mqtp.QueueRoute {
	return mqtp.QueueRoute("task_assignment")
}
