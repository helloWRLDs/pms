package list

type Filters struct {
	Pagination
	Created
	OrderBy string
}

type Created struct {
	CreatedFrom string `json:"created_from"`
	CreatedTo   string `json:"created_to"`
}
