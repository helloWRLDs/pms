package ws

type Action[T any] struct {
	Action string `json:"action"`
	ID     string `json:"id"`
	Value  T      `json:"value"`
}
