package router

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"pms.api-gateway/internal/models"
	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/transport/ws"
)

func (s *Server) StreamDocument(c *websocket.Conn) {
	log := s.log.Named("StreamDocument").With(
		zap.String("ip", c.IP()),
	)
	log.Debug("StreamDocument called")
	log.Info("client connected")

	docID := c.Params("docID")
	if strings.Trim(docID, " ") == "" {
		return
	}
	_, err := s.Logic.GetDocument(context.Background(), docID)
	if err != nil {
		return
	}

	hubID := fmt.Sprintf("doc-%s", docID)
	if _, exists := s.Logic.WsHubs[hubID]; !exists {
		s.Logic.WsHubs[hubID] = ws.NewHub()
		doc, err := s.Logic.GetDocument(context.Background(), docID)
		if err != nil {
			log.Errorw("failed to find doc", "err", err)
			return
		}
		s.Logic.DocumentsCache.Set(context.Background(), hubID, models.DocumentBody{
			RequireUpdate: false,
			Document:      doc,
		}, 24)
	}

	s.Logic.WsHubs[hubID].AddClient(c)
	defer func() {
		s.Logic.WsHubs[hubID].RemoveClient(c)
		c.Close()
	}()

	docBody, _ := s.Logic.DocumentsCache.Get(context.Background(), hubID)

	if raw, err := json.Marshal(docBody.Document); err != nil {
		log.Errorw("failed marshaling doc body", "err", err)
		return
	} else {
		s.Logic.WsHubs[hubID].Broadcast(raw)
	}

	var (
		mt  int
		msg []byte
	)

	for {
		mt, msg, err = c.ReadMessage()
		if err != nil {
			log.Infow("read:", err)
			break
		}
		log.Infof("recv: %s mt: %d", msg, mt)

		if msg == nil {
			continue
		}
		doc := new(dto.Document)
		if err := json.Unmarshal(msg, doc); err != nil {
			log.Errorw("failed to resolve response")
		} else {
			s.Logic.DocumentsCache.Set(context.Background(), hubID, models.DocumentBody{
				Document:      doc,
				RequireUpdate: true,
			}, 24)
		}
	}
}

func (s *Server) StreamSprint(c *websocket.Conn) {
	log := s.log.Named("StreamSprint").With(
		zap.String("ip", c.IP()),
	)
	log.Debug("StreamSprint called")
	log.Info("client connected")

	sprintID := c.Params("sprintID")
	if strings.Trim(sprintID, " ") == "" {
		log.Error("sprint_id is invalid")
		return
	}
	hubID := fmt.Sprintf("sprint-%s", sprintID)
	defer func() {
		s.Logic.WsHubs[hubID].RemoveClient(c)
		c.Close()
	}()

	if _, exists := s.Logic.WsHubs[hubID]; !exists {
		log.Info("hub not found. Creating...")
		s.Logic.WsHubs[hubID] = ws.NewHub()
	}
	s.Logic.WsHubs[hubID].AddClient(c)

	var (
		mt  int
		msg []byte
		err error
	)

	log.Infow("trying to list tasks", "sprint_id", sprintID)
	tasks, err := s.Logic.ListTasks(context.Background(), &dto.TaskFilter{
		SprintId: sprintID,
		// AssigneeId: c.Query("assignee_id"),
		Page:    1,
		PerPage: 10000,
	})
	if err != nil {
		log.Errorw("failed to fetch tasks", "err", err)
		return
	} else {
		log.Infow("fetched tasks", "tasks", tasks)
		c.WriteJSON(tasks.Items)
	}

	for {
		mt, msg, err = c.ReadMessage()
		if err != nil {
			log.Infow("read:", err)
			break
		}
		log.Infof("recv: %s mt: %d", msg, mt)

		if msg != nil {
			task := new(dto.Task)
			if err := json.Unmarshal(msg, task); err != nil {
				log.Errorw("failed to resolve response")
			} else {
				s.Logic.TaskQueue.Rpush(context.Background(), hubID, models.TaskQueueElement{
					Value: task,
				})
			}
		}
	}
}

func (s *Server) UpdateSprint(c *fiber.Ctx) error {
	log := s.log.Named("UpdateSprint").With(
		zap.String("ip", c.IP()),
	)
	log.Debug("UpdateSprint called")

	sprintID := c.Params("sprintID", "")
	if strings.Trim(sprintID, " ") == "" {
		return errs.ErrBadGateway{
			Object: "sprint_id",
		}
	}
	sprint := new(dto.Sprint)
	if err := c.BodyParser(sprint); err != nil {
		return err
	}

	updated, err := s.Logic.UpdateSprint(c.UserContext(), sprintID, sprint)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(updated)
}

func (s *Server) CreateSprint(c *fiber.Ctx) error {
	log := s.log.Named("CreateSprint").With(
		zap.String("ip", c.IP()),
	)
	log.Debug("CreateSprint called")

	creation := new(dto.SprintCreation)
	if err := c.BodyParser(creation); err != nil {
		return err
	}
	creation.ProjectId = c.Locals("project_id").(string)

	created, err := s.Logic.CreateSprint(c.UserContext(), creation)
	if err != nil {
		return err
	}
	return c.Status(201).JSON(created)
}

func (s *Server) ListSprints(c *fiber.Ctx) error {
	log := s.log.Named("ListSprints").With(
		zap.String("ip", c.IP()),
	)
	log.Debug("ListCompanies called")

	filter := &dto.SprintFilter{
		Page:      int32(c.QueryInt("page", 1)),
		PerPage:   int32(c.QueryInt("per_page", 10)),
		ProjectId: c.Locals("project_id").(string),
		Title:     c.Query("title", ""),
	}
	sprints, err := s.Logic.ListSprints(c.UserContext(), filter)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(sprints)
}

func (s *Server) GetSprint(c *fiber.Ctx) error {
	log := s.log.Named("GetSprint").With(
		zap.String("ip", c.IP()),
	)
	log.Debug("GetSprint called")

	sprintID := c.Params("sprintID", "")
	if strings.Trim(sprintID, " ") == "" {
		return errs.ErrBadGateway{
			Object: "sprint_id",
		}
	}

	sprint, err := s.Logic.GetSprint(c.UserContext(), sprintID)
	if err != nil {
		return err
	}
	if sprint.ProjectId != c.Locals("project_id").(string) {
		return errs.ErrBadGateway{
			Object: "sprint_id",
		}
	}

	return c.Status(200).JSON(sprint)
}
