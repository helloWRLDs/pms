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

func (s *Server) StreamSprint(c *websocket.Conn) {
	log := s.log.Named("StreamSprint").With(
		zap.String("ip", c.IP()),
	)
	log.Debug("StreamSprint called")

	sprintID := c.Params("sprintID")
	if strings.Trim(sprintID, " ") == "" {
		return
	}
	hubID := fmt.Sprintf("sprint-%s", sprintID)
	defer func() {
		s.wshubs[hubID].RemoveClient(c)
		c.Close()
	}()

	if _, exists := s.wshubs[hubID]; !exists {
		s.wshubs[hubID] = ws.NewHub()
	}
	s.wshubs[hubID].AddClient(c)

	sendTasks := func(c *websocket.Conn) error {
		tasks, err := s.Logic.ListTasks(context.Background(), &dto.TaskFilter{
			SprintId:   sprintID,
			AssigneeId: c.Query("assignee_id"),
			Page:       1,
			PerPage:    10000,
		})
		if err != nil {
			log.Errorw("failed to fetch tasks", "err", err)
			return err
		} else {
			log.Infow("fetched tasks", "tasks", tasks)
			c.WriteJSON(tasks.Items)
		}
		return nil
	}

	sendTasks(c)
	var (
		mt  int
		msg []byte
		err error
	)

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

		if err := sendTasks(c); err != nil {
			log.Errorw("failed to send tasks", "err", err)
		}

		// if err = c.WriteMessage(mt, msg); err != nil {
		// 	log.Infow("write:", err)
		// 	break
		// }
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
