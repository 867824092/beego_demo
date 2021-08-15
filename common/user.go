//package usermodel
package common

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewUser(id int, name string, age int) *User {
	user := &User{
		id,
		name,
		age,
	}
	return user
}
