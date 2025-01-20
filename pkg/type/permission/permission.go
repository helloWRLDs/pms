package permission

type Permission uint16

const (
	ORG_READ Permission = iota
	ORG_UPDATE
	ORG_MANAGE

	USER_READ
)

func ParsePermission(name string) Permission {
	switch name {
	case "ORG_READ":
		return ORG_READ
	case "ORG_UPDATE":
		return ORG_UPDATE
	case "ORG_MANAGE":
		return ORG_MANAGE
	case "USER_READ":
		return USER_READ
	default:
		return 0
	}
}

func (p Permission) String() string {
	switch p {
	case USER_READ:
		return "USER_READ"
	case ORG_READ:
		return "ORG_READ"
	case ORG_UPDATE:
		return "ORG_UPDATE"
	case ORG_MANAGE:
		return "ORG_MANAGE"
	default:
		return ""
	}
}
