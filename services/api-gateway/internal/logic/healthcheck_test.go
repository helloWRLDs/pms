package logic

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"pms.api-gateway/internal/config"
	"pms.pkg/datastore/redis"
	"pms.pkg/logger"
	"pms.pkg/tools/scheduler"
	configgrpc "pms.pkg/transport/grpc/config"
)

func TestIntegration_HealthChecks(t *testing.T) {

	cfg := config.Config{
		Auth: configgrpc.ClientConfig{
			Host: "localhost:50051",
		},
		Project: configgrpc.ClientConfig{
			Host: "localhost:50052",
		},
		Analytics: configgrpc.ClientConfig{
			Host: "localhost:50053",
		},
		Redis: redis.Config{
			Host: "localhost:6379",
		},
	}
	log := logger.Log

	l := &Logic{
		Config: cfg,
		log:    log,
		Tasks:  make(map[string]*scheduler.Task),
	}

	l.InitTasks()

	t.Run("Check Auth Health", func(t *testing.T) {
		err := l.CheckAuthHealth(context.Background())
		require.NoError(t, err)
		assert.NotNil(t, l.authClient)
		assert.True(t, l.authClient.Ping())
	})

	t.Run("Check Project Health", func(t *testing.T) {
		err := l.CheckProjectHealth(context.Background())
		require.NoError(t, err)
		assert.NotNil(t, l.projectClient)
		assert.True(t, l.projectClient.Ping())
	})

	t.Run("Check Analytics Health", func(t *testing.T) {
		err := l.CheckAnalyticsHealth(context.Background())
		require.NoError(t, err)
		assert.NotNil(t, l.analyticsClient)
		assert.True(t, l.analyticsClient.Ping())
	})

	t.Run("Check Notifier Health", func(t *testing.T) {
		err := l.CheckNotifierHealth(context.Background())
		require.NoError(t, err)
		assert.NotNil(t, l.notificationMQ)
		assert.Equal(t, "READY", l.notificationMQ.ConnState())
	})

	t.Run("Close Clients", func(t *testing.T) {
		require.NoError(t, l.CheckAuthHealth(context.Background()))
		require.NoError(t, l.CheckProjectHealth(context.Background()))
		require.NoError(t, l.CheckAnalyticsHealth(context.Background()))
		require.NoError(t, l.CheckNotifierHealth(context.Background()))

		l.CloseClients(context.Background())

		assert.False(t, l.authClient.Ping())
		assert.False(t, l.projectClient.Ping())
		assert.False(t, l.analyticsClient.Ping())
		assert.NotEqual(t, "READY", l.notificationMQ.ConnState())
	})

	t.Run("Health Check with Invalid Hosts", func(t *testing.T) {
		originalAuthHost := cfg.Auth.Host
		cfg.Auth.Host = "invalid:12345"
		l.Config = cfg

		err := l.CheckAuthHealth(context.Background())
		assert.Error(t, err)

		cfg.Auth.Host = originalAuthHost
		l.Config = cfg
	})

	t.Run("Concurrent Health Checks", func(t *testing.T) {
		done := make(chan bool)
		for i := 0; i < 5; i++ {
			go func() {
				require.NoError(t, l.CheckAuthHealth(context.Background()))
				require.NoError(t, l.CheckProjectHealth(context.Background()))
				require.NoError(t, l.CheckAnalyticsHealth(context.Background()))
				require.NoError(t, l.CheckNotifierHealth(context.Background()))
				done <- true
			}()
		}

		for i := 0; i < 5; i++ {
			select {
			case <-done:
				continue
			case <-time.After(5 * time.Second):
				t.Fatal("Timeout waiting for concurrent health checks")
			}
		}
	})
}
