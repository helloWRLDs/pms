package logic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"pms.api-gateway/internal/config"
	"pms.pkg/logger"
	configgrpc "pms.pkg/transport/grpc/config"
)

func TestIntegration_Analytics(t *testing.T) {

	cfg := config.Config{
		Analytics: configgrpc.ClientConfig{
			Host: "localhost:50053",
		},
	}
	log := logger.Log

	l := &Logic{
		Config: cfg,
		log:    log,
	}

	t.Run("Get Project Stats", func(t *testing.T) {
		companyID := "test-company-1"
		stats, err := l.GetProjectStats(context.Background(), companyID)
		require.NoError(t, err)
		assert.NotNil(t, stats)

		for _, stat := range stats {
			assert.NotEmpty(t, stat.UserId)
			assert.NotEmpty(t, stat.FirstName)
			assert.NotEmpty(t, stat.LastName)
			assert.NotNil(t, stat.Stats)
			for _, taskStats := range stat.Stats {
				assert.GreaterOrEqual(t, taskStats.TotalTasks, int32(0))
				assert.GreaterOrEqual(t, taskStats.DoneTasks, int32(0))
				assert.GreaterOrEqual(t, taskStats.InProgressTasks, int32(0))
				assert.GreaterOrEqual(t, taskStats.ToDoTasks, int32(0))
				assert.GreaterOrEqual(t, taskStats.TotalPoints, int32(0))
			}
		}
	})

	t.Run("Get Project Stats with Invalid Company ID", func(t *testing.T) {
		stats, err := l.GetProjectStats(context.Background(), "non-existent-company")
		assert.Error(t, err)
		assert.Nil(t, stats)
	})

	t.Run("Get Project Stats with Empty Company ID", func(t *testing.T) {
		stats, err := l.GetProjectStats(context.Background(), "")
		assert.Error(t, err)
		assert.Nil(t, stats)
	})

	t.Run("Get Project Stats with Analytics Service Unavailable", func(t *testing.T) {
		originalHost := cfg.Analytics.Host
		cfg.Analytics.Host = "invalid:12345"
		l.Config = cfg

		stats, err := l.GetProjectStats(context.Background(), "test-company-1")
		assert.Error(t, err)
		assert.Nil(t, stats)

		cfg.Analytics.Host = originalHost
		l.Config = cfg
	})
}
