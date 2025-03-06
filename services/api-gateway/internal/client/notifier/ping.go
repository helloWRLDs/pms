package notifierclient

import "google.golang.org/grpc/connectivity"

func (c *NotifierClient) Ping() bool {
	return c.conn.GetState() == connectivity.Ready
}
