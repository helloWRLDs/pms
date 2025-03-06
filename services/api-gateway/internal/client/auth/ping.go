package authclient

import "google.golang.org/grpc/connectivity"

func (c *AuthClient) Ping() bool {
	c.log.Infof("Current connection state: %s", c.State())

	if c.State() == connectivity.Idle {
		c.log.Warn("Connection is IDLE, attempting to wake it up...")
		c.conn.Connect() // Forces transition from IDLE to CONNECTING/READY
	}

	return c.State() == connectivity.Ready
}

func (c *AuthClient) State() connectivity.State {
	return c.conn.GetState()
}
