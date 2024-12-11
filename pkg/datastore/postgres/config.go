package postgres

import (
	"fmt"
	"strings"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSL      string
}

func (c Config) DSN() string {
	dsn := []string{
		fmt.Sprintf("user=%s", c.User),
		fmt.Sprintf("password=%s", c.Password),
		fmt.Sprintf("dbname=%s", c.Name),
	}
	return strings.Join(dsn, " ")
}
