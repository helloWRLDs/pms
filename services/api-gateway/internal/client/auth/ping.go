package authclient

import "google.golang.org/grpc/connectivity"

func (c *AuthClient) Ping() bool {
	state := c.conn.GetState()
	c.log.Infof("Current connection state: %s", state)

	if state == connectivity.Idle {
		c.log.Warn("Connection is IDLE, attempting to wake it up...")
		c.conn.Connect() // Forces transition from IDLE to CONNECTING/READY
	}

	return c.conn.GetState() == connectivity.Ready
}
