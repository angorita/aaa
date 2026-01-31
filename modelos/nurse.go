package modelos

import (
	"time"
)

type Nurse struct {
	Id     int
	Nombre string
	Hoy    time.Time
	Estado bool
}

func (n *Nurse) AddNurse(id int, name string, createAdd time.Time, status bool) {
	n.Id = id
	n.Nombre = name
	n.Hoy = createAdd
	n.Estado = status
}
