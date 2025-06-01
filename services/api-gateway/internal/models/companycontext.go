package models

type Participant struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type CompanyContext struct {
	CompanyID    string        `json:"company_id"`
	Sprints      []string      `json:"sprints"`
	Projects     []string      `json:"projects"`
	Participants []Participant `json:"participants"`
}

func (c CompanyContext) GetDB() int {
	return 3
}
