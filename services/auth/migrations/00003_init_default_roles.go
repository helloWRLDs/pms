package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
	"pms.auth/internal/entity"
	"pms.pkg/consts"
)

func init() {
	goose.AddMigrationContext(upInitDefaultRoles, downInitDefaultRoles)
}

func upInitDefaultRoles(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	roles := []entity.Role{
		{
			Name: "admin",
			Persmissions: consts.PermissionSet{
				consts.ORG_READ_PERMISSION,
				consts.ORG_WRITE_PERMISSION,
				consts.USER_DELETE_PERMISSION,
				consts.USER_READ_PERMISSION,
				consts.USER_WRITE_PERMISSION,
				consts.PROJECT_READ_PERMISSION,
				consts.PROJECT_WRITE_PERMISSION,
				consts.PROJECT_DELETE_PERMISSION,
				consts.PROJECT_ADD_PERMISSION,
				consts.TASK_READ_PERMISSION,
				consts.TASK_WRITE_PERMISSION,
				consts.TASK_DELETE_PERMISSION,
				consts.TASK_ADD_PERMISSION,
				consts.ROLE_READ_PERMISSION,
				consts.ROLE_WRITE_PERMISSION,
				consts.ROLE_DELETE_PERMISSION,
				consts.ROLE_ADD_PERMISSION,
			},
		},
		{
			Name: "reader",
			Persmissions: consts.PermissionSet{
				consts.PROJECT_READ_PERMISSION,
				consts.TASK_READ_PERMISSION,
				consts.ORG_READ_PERMISSION,
			},
		},
	}
	q := `
	INSERT INTO Roles(name, permissions) VALUES (?, ?);
	`
	for _, role := range roles {
		_, err := tx.Exec(q, role.Name, role.Persmissions)
		if err != nil {
			return err
		}
	}
	return nil
}

func downInitDefaultRoles(ctx context.Context, tx *sql.Tx) error {
	if _, err := tx.Exec("DELETE FROM Roles WHERE name in ('admin', 'reader')"); err != nil {
		return err
	}
	return nil
}
