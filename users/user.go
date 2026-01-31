package users

import (
	"fmt"
	"time"

	"github.com/angorita/aaa/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Mengeche", time.Now(), true)
	fmt.Print(u)
}
func Nurse() {
	n := new(modelos.Nurse)
	n.AddNurse(10, "Ariel", time.Now(), true)
	fmt.Print(n)
}
