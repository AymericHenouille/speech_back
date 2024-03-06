package databases

import "fmt"

type OpenDBInfo struct {
	User     string
	Password string
	Host     string
	Name     string
}

func New(user, password, host, name string) *OpenDBInfo {
	return &OpenDBInfo{
		User:     user,
		Password: password,
		Host:     host,
		Name:     name,
	}
}

func (o *OpenDBInfo) Datasource() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", o.User, o.Password, o.Host, o.Name)
}
