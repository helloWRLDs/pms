package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
	"pms.auth/internal/entity"
	"pms.pkg/type/permissions"
)

func init() {
	goose.AddMigrationContext(upInitDefaultRoles, downInitDefaultRoles)
}

func upInitDefaultRoles(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	roles := []entity.Role{
		{
			Name: "admin",
			Persmissions: permissions.PermissionSet{
				permissions.ORG_READ_PERMISSION,
				permissions.ORG_WRITE_PERMISSION,
				permissions.USER_DELETE_PERMISSION,
				permissions.USER_READ_PERMISSION,
				permissions.USER_WRITE_PERMISSION,
				permissions.PROJECT_READ_PERMISSION,
				permissions.PROJECT_WRITE_PERMISSION,
				permissions.PROJECT_DELETE_PERMISSION,
				permissions.PROJECT_ADD_PERMISSION,
				permissions.TASK_READ_PERMISSION,
				permissions.TASK_WRITE_PERMISSION,
				permissions.TASK_DELETE_PERMISSION,
				permissions.TASK_ADD_PERMISSION,
				permissions.ROLE_READ_PERMISSION,
				permissions.ROLE_WRITE_PERMISSION,
				permissions.ROLE_DELETE_PERMISSION,
				permissions.ROLE_ADD_PERMISSION,
			},
		},
		{
			Name: "reader",
			Persmissions: permissions.PermissionSet{
				permissions.PROJECT_READ_PERMISSION,
				permissions.TASK_READ_PERMISSION,
				permissions.ORG_READ_PERMISSION,
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
