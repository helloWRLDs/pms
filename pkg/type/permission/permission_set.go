package permission

type PermissionSet []Permission

func (p PermissionSet) Array() []Permission {
	return p
}

func (p PermissionSet) StringArray() []string {
	arr := make([]string, len(p))
	for i, p := range p {
		arr[i] = p.String()
	}
	return arr
}
