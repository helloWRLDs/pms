package data

import (
	"github.com/lib/pq"
	"go.uber.org/zap"
	"pms.pkg/consts"
	"pms.pkg/logger"
)

func (r *Repository) MigrateUp() {
	log := logger.Log.Named("MigrateUp")
	log.Debugw("migrating up")

	err := r.MigrateAdminRole()
	if err != nil {
		log.Errorw("failed to migrate admin role", "err", err)
		return
	}
	log.Debugw("migrated admin role")
}

func (r *Repository) MigrateAdminRole() error {
	log := logger.Log.With(zap.String("func", "MigrateAdminRole"))
	log.Debugw("migrating admin role")

	permissionSet := []consts.Permission{
		consts.COMPANY_READ_PERMISSION,
		consts.COMPANY_WRITE_PERMISSION,
		consts.COMPANY_DELETE_PERMISSION,
		consts.COMPANY_INVITE_PERMISSION,
		consts.USER_READ_PERMISSION,
		consts.USER_WRITE_PERMISSION,
		consts.USER_DELETE_PERMISSION,
		consts.USER_INVITE_PERMISSION,
		consts.PROJECT_READ_PERMISSION,
		consts.PROJECT_WRITE_PERMISSION,
		consts.PROJECT_DELETE_PERMISSION,
		consts.PROJECT_INVITE_PERMISSION,
		consts.TASK_READ_PERMISSION,
		consts.TASK_WRITE_PERMISSION,
		consts.TASK_DELETE_PERMISSION,
		consts.TASK_INVITE_PERMISSION,
		consts.ROLE_READ_PERMISSION,
		consts.ROLE_WRITE_PERMISSION,
		consts.ROLE_DELETE_PERMISSION,
		consts.ROLE_INVITE_PERMISSION,
		consts.SPRINT_READ_PERMISSION,
		consts.SPRINT_WRITE_PERMISSION,
		consts.SPRINT_DELETE_PERMISSION,
		consts.SPRINT_INVITE_PERMISSION,
	}

	q := `
	INSERT INTO "Role" (name, permissions) 
	VALUES ($1, $2) 
	ON CONFLICT DO NOTHING
	`
	_, err := r.db.Exec(q, "admin", pq.Array(permissionSet))
	if err != nil {
		log.Errorw("failed to migrate default roles", "err", err)
		return err
	}
	log.Debugw("migrated default roles")
	return nil
}
