package nullable

import "database/sql"

func String(value string) (str sql.NullString) {
	str.String = value
	if value == "" {
		str.Valid = false
	} else {
		str.Valid = true
	}
	return
}
