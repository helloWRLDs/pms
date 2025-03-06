package client

type Client interface {
	Close() error
	Ping() bool
}
