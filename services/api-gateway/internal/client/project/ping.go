package projectclient

import "google.golang.org/grpc/connectivity"

func (c *ProjectClient) Ping() bool {
	c.log.Debugf("Current connection state: %s", c.State())

	if c.State() == connectivity.Idle {
		c.log.Warn("Connection is IDLE, attempting to wake it up...")
		c.conn.Connect()
	}

	return c.State() == connectivity.Ready
}

func (c *ProjectClient) State() connectivity.State {
	return c.conn.GetState()
}
