package logic

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"pms.api-gateway/internal/config"
	"pms.pkg/datastore/redis"
	"pms.pkg/logger"
	configgrpc "pms.pkg/transport/grpc/config"
	"pms.pkg/transport/ws"
)

func TestIntegration_StreamProcessing(t *testing.T) {

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
		WsHubs: make(map[string]*ws.Hub),
	}

	docID := uuid.New().String()
	docHubID := "doc-" + docID
	docHub := ws.NewHub()
	l.WsHubs[docHubID] = docHub

	sprintID := uuid.New().String()
	sprintHubID := "sprint-" + sprintID
	sprintHub := ws.NewHub()
	l.WsHubs[sprintHubID] = sprintHub

	t.Run("Process Document Stream", func(t *testing.T) {
		go l.processDocumentStream()

		time.Sleep(2 * time.Second)

		assert.NotNil(t, l.WsHubs[docHubID])
		assert.Equal(t, docHub, l.WsHubs[docHubID])

		invalidHubID := "invalid-hub"
		l.WsHubs[invalidHubID] = ws.NewHub()
		time.Sleep(2 * time.Second)
		assert.NotNil(t, l.WsHubs[invalidHubID])
	})

	t.Run("Process Task Stream", func(t *testing.T) {
		go l.processTaskStream()

		time.Sleep(6 * time.Second)

		assert.NotNil(t, l.WsHubs[sprintHubID])
		assert.Equal(t, sprintHub, l.WsHubs[sprintHubID])

		invalidHubID := "invalid-hub"
		l.WsHubs[invalidHubID] = ws.NewHub()
		time.Sleep(6 * time.Second)
		assert.NotNil(t, l.WsHubs[invalidHubID])
	})

	t.Run("Concurrent Stream Processing", func(t *testing.T) {
		hubs := make(map[string]*ws.Hub)
		for i := 0; i < 5; i++ {
			docID := uuid.New().String()
			hubID := "doc-" + docID
			hubs[hubID] = ws.NewHub()
			l.WsHubs[hubID] = hubs[hubID]

			sprintID := uuid.New().String()
			hubID = "sprint-" + sprintID
			hubs[hubID] = ws.NewHub()
			l.WsHubs[hubID] = hubs[hubID]
		}

		go l.processDocumentStream()
		go l.processTaskStream()

		time.Sleep(6 * time.Second)

		for hubID, hub := range hubs {
			assert.NotNil(t, l.WsHubs[hubID])
			assert.Equal(t, hub, l.WsHubs[hubID])
		}
	})

	t.Run("Hub Cleanup", func(t *testing.T) {
		sprintID := uuid.New().String()
		hubID := "sprint-" + sprintID
		hub := ws.NewHub()
		l.WsHubs[hubID] = hub

		go l.processTaskStream()

		time.Sleep(6 * time.Second)

		assert.NotContains(t, l.WsHubs, hubID)
	})
}
