package router

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
)

func (s *Server) DownloadDocument(c *fiber.Ctx) error {
	log := s.log.Named("DownloadDocument").With(
		zap.String("ip", c.IP()),
	)
	log.Info("DownloadDocument called")

	docID := c.Params("docID", "")
	c.Set("Content-Type", "application/pdf")

	docPDF, err := s.Logic.DownloadDocumentPDF(c.UserContext(), docID)
	if err != nil {
		return err
	}
	c.Set("X-Document-Title", docPDF.Title)
	c.Set("X-Document-ID", docPDF.DocId)
	c.Set("Content-Type", "application/pdf")

	return c.Status(200).Send(docPDF.Body)
}

func (s *Server) CreateReportTemplate(c *fiber.Ctx) error {
	log := s.log.Named("CreateReportTemplate").With(
		zap.String("ip", c.IP()),
	)
	log.Info("CreateReportTemplate called")

	docCreation := new(dto.DocumentCreation)

	if err := c.BodyParser(docCreation); err != nil {
		log.Error("failed to resolve sprintID")
		return errs.ErrBadGateway{
			Object: "sprint_id",
		}
	}

	docID, err := s.Logic.CreateReportTemplate(c.UserContext(), docCreation)
	if err != nil {
		log.Errorw("faield to create template", "err", err)
		return err
	}
	return c.Status(200).JSON(map[string]string{
		"msg": fmt.Sprintf("created doc with id = %s", docID),
	})
}

func (s *Server) GetDocument(c *fiber.Ctx) error {
	log := s.log.Named("GetDocument").With(
		zap.String("ip", c.IP()),
	)
	log.Info("GetDocument called")

	docId := c.Params("docID", "")
	if strings.Trim(docId, " ") == "" {
		return errs.ErrBadGateway{
			Object: "doc_id",
		}
	}

	doc, err := s.Logic.GetDocument(c.UserContext(), docId)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(doc)
}

func (s *Server) ListDocuments(c *fiber.Ctx) error {
	log := s.log.Named("ListDocuments").With(
		zap.String("ip", c.IP()),
	)
	log.Info("ListDocuments called")

	filter := &dto.DocumentFilter{
		Page:      int32(c.QueryInt("page", 1)),
		PerPage:   int32(c.QueryInt("per_page", 10)),
		Title:     c.Query("title", ""),
		ProjectId: c.Query("project_id", ""),
	}
	docs, err := s.Logic.ListDocuments(c.UserContext(), filter)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(docs)
}

func (s *Server) UpdateDocument(c *fiber.Ctx) error {
	log := s.log.Named("UpdateDocument").With(
		zap.String("ip", c.IP()),
	)
	log.Info("UpdateDocument called")

	docId := c.Params("docID", "")
	if strings.Trim(docId, " ") == "" {
		return errs.ErrBadGateway{
			Object: "doc_id",
		}
	}

	updatedDoc := new(dto.Document)
	if err := c.BodyParser(updatedDoc); err != nil {
		return errs.ErrBadGateway{
			Object: "doc body",
		}
	}
	if err := s.Logic.UpdateDocument(c.UserContext(), docId, updatedDoc); err != nil {
		return err
	}
	return c.SendStatus(200)
}
