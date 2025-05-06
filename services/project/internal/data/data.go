package data

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	assignmentdata "pms.project/internal/data/assignment"
	projectdata "pms.project/internal/data/project"
	sprintdata "pms.project/internal/data/sprint"
	taskdata "pms.project/internal/data/task"
)

type Repository struct {
	db *sqlx.DB

	Project        *projectdata.Repository
	Sprint         *sprintdata.Repository
	Task           *taskdata.Repository
	TaskAssignment *assignmentdata.Repository
}

func New(db *sqlx.DB, log *zap.SugaredLogger) *Repository {
	return &Repository{
		db:             db,
		Project:        projectdata.New(db, log),
		Task:           taskdata.New(db, log),
		Sprint:         sprintdata.New(db, log),
		TaskAssignment: assignmentdata.New(db, log),
	}
}
