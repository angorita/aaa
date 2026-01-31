package modelos

import (
	"time"
)

type User struct {
	Id        int
	Name      string
	CreateAdd time.Time
	Status    bool
}

func (u *User) AddUser(id int, name string, createAdd time.Time, status bool) {
	u.Id = id
	u.Name = name
	u.CreateAdd = createAdd
	u.Status = status
}
