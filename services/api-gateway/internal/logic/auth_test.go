package logic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	authclient "pms.api-gateway/internal/client/auth"
	"pms.api-gateway/internal/config"
	"pms.api-gateway/internal/models"
	"pms.pkg/datastore/redis"
	"pms.pkg/logger"
	configgrpc "pms.pkg/transport/grpc/config"
	"pms.pkg/transport/grpc/dto"
)

func TestIntegration_AuthFlow(t *testing.T) {

	cfg := config.Config{
		Auth: configgrpc.ClientConfig{
			Host: "localhost:50051",
		},
		Redis: redis.Config{
			Host: "localhost:6379",
		},
	}
	log := logger.Log

	authClient, err := authclient.New(cfg.Auth, log)
	require.NoError(t, err)
	defer authClient.Close()

	l := &Logic{
		Config:     cfg,
		authClient: authClient,
		Sessions:   redis.New(&cfg.Redis, models.Session{}),
		log:        log,
	}

	t.Run("Register and Login Flow", func(t *testing.T) {
		newUser := &dto.NewUser{
			Email:     "test@example.com",
			Password:  "password123",
			FirstName: "Test",
			LastName:  "User",
		}
		user, err := l.RegisterUser(context.Background(), newUser)
		require.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, newUser.Email, user.Email)
		assert.Equal(t, newUser.FirstName, user.FirstName)
		assert.Equal(t, newUser.LastName, user.LastName)

		creds := &dto.UserCredentials{
			Email:    newUser.Email,
			Password: newUser.Password,
		}
		payload, err := l.LoginUser(context.Background(), creds)
		require.NoError(t, err)
		assert.NotNil(t, payload)
		assert.NotEmpty(t, payload.SessionId)
		assert.NotEmpty(t, payload.AccessToken)
		assert.NotEmpty(t, payload.RefreshToken)
		assert.NotNil(t, payload.User)
		assert.Equal(t, user.Id, payload.User.Id)
	})

	t.Run("OAuth2 Flow", func(t *testing.T) {
		provider := "google"
		authURL, err := l.InitiateOAuth2(context.Background(), provider)
		require.NoError(t, err)
		assert.NotEmpty(t, authURL)
		assert.Contains(t, authURL, "accounts.google.com")

	})
}

func TestIntegration_InvalidCredentials(t *testing.T) {
	cfg := config.Config{
		Auth: configgrpc.ClientConfig{
			Host: "localhost:50051",
		},
		Redis: redis.Config{
			Host: "localhost:6379",
		},
	}
	log := logger.Log

	authClient, err := authclient.New(cfg.Auth, log)
	require.NoError(t, err)
	defer authClient.Close()

	l := &Logic{
		Config:     cfg,
		authClient: authClient,
		Sessions:   redis.New(&cfg.Redis, models.Session{}),
		log:        log,
	}

	t.Run("Login with Invalid Credentials", func(t *testing.T) {
		creds := &dto.UserCredentials{
			Email:    "nonexistent@example.com",
			Password: "wrongpassword",
		}
		payload, err := l.LoginUser(context.Background(), creds)
		assert.Error(t, err)
		assert.Nil(t, payload)
	})

	t.Run("Register with Invalid Data", func(t *testing.T) {
		invalidUser := &dto.NewUser{
			Email:     "invalid-email",
			Password:  "password123",
			FirstName: "Test",
			LastName:  "User",
		}
		user, err := l.RegisterUser(context.Background(), invalidUser)
		assert.Error(t, err)
		assert.Nil(t, user)

		emptyPassUser := &dto.NewUser{
			Email:     "test2@example.com",
			Password:  "",
			FirstName: "Test",
			LastName:  "User",
		}
		user, err = l.RegisterUser(context.Background(), emptyPassUser)
		assert.Error(t, err)
		assert.Nil(t, user)
	})
}
