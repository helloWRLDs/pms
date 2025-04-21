package data

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	companydata "pms.auth/internal/data/company"
	participantdata "pms.auth/internal/data/participant"
	roledata "pms.auth/internal/data/role"
	userdata "pms.auth/internal/data/user"
)

type Repository struct {
	db *sqlx.DB

	User        *userdata.Repository
	Company     *companydata.Repository
	Role        *roledata.Repository
	Participant *participantdata.Repository
}

func New(db *sqlx.DB, log *zap.SugaredLogger) *Repository {
	return &Repository{
		db:          db,
		User:        userdata.New(db, log),
		Company:     companydata.New(db, log),
		Role:        roledata.New(db, log),
		Participant: participantdata.New(db, log),
	}
}
