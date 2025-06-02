package logic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	authclient "pms.api-gateway/internal/client/auth"
	projectclient "pms.api-gateway/internal/client/project"
	"pms.api-gateway/internal/config"
	"pms.api-gateway/internal/models"
	"pms.pkg/datastore/redis"
	"pms.pkg/logger"
	configgrpc "pms.pkg/transport/grpc/config"
)

func TestIntegration_CompanyContext(t *testing.T) {

	cfg := config.Config{
		Auth: configgrpc.ClientConfig{
			Host: "localhost:50051",
		},
		Project: configgrpc.ClientConfig{
			Host: "localhost:50052",
		},
		Redis: redis.Config{
			Host: "localhost:6379",
		},
	}
	log := logger.Log

	authClient, err := authclient.New(cfg.Auth, log)
	require.NoError(t, err)
	defer authClient.Close()

	projectClient, err := projectclient.New(cfg.Project, log)
	require.NoError(t, err)
	defer projectClient.Close()

	l := &Logic{
		Config:         cfg,
		authClient:     authClient,
		projectClient:  projectClient,
		CompanyContext: redis.New(&cfg.Redis, models.CompanyContext{}),
		log:            log,
	}

	t.Run("Get Company Context", func(t *testing.T) {

		companyID := "test-company-1" // real company id

		cc, err := l.GetCompanyContext(context.Background(), companyID)
		require.NoError(t, err)
		assert.NotNil(t, cc)
		assert.Equal(t, companyID, cc.CompanyID)

		assert.NotNil(t, cc.Projects)

		assert.NotNil(t, cc.Sprints)

		assert.NotNil(t, cc.Participants)
		for _, p := range cc.Participants {
			assert.NotEmpty(t, p.UserID)
			assert.NotEmpty(t, p.FirstName)
			assert.NotEmpty(t, p.LastName)
			assert.NotEmpty(t, p.Email)
		}
	})

	t.Run("Get Company Context with Invalid ID", func(t *testing.T) {
		cc, err := l.GetCompanyContext(context.Background(), "non-existent-company")
		assert.Error(t, err)
		assert.Nil(t, cc)
	})

	t.Run("Get Company Context with Empty ID", func(t *testing.T) {
		cc, err := l.GetCompanyContext(context.Background(), "")
		assert.Error(t, err)
		assert.Nil(t, cc)
	})
}
