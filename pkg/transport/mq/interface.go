package mqtp

type Queueable interface {
	RoutingKey() QueueRoute
}

type QueueRoute string

type QueueName string
