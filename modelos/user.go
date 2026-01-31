package modelos

import(
    "time"
)

type User struct{
    Id int
    Name string
    Create time.Time
    Status bool
}

func(this User)AddUser(id int,name string,createAdd time.Time,status bool){
    this.Id=id
    this.Name=name
    this.CreateAdd=createAdd
    this.Status=status
}

