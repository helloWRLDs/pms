package roledomain

import "pms.pkg/type/permission"

type Role struct {
	ID             int                      `db:"id"`
	Name           string                   `db:"name"`
	Permissions    permission.PermissionSet `db:"permissions"`
	OrganizationID string                   `db:"org_id"`
}
