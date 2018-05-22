package _struct

import "fmt"

type User struct {
	Id    int
	Name  string
	Bio   string
	Email string
}

type Student struct {
	Id    int    `validate:"number,min=5,max=1000"`
	Name  string `validate:"string,min=2,max=10"`
	Bio   string `validate:"string"`
	Email string `validate:"email"`
}

func main() {
	user := User{
		Id:    0,
		Name:  "ali",
		Bio:   "a",
		Email: "b",
	}
	checkUserId(user)

	student := Student{
		Id:    0,
		Name:  "fdas",
		Bio:   "fdsa",
		Email: "fdsafdsa",
	}
	fmt.Println(student.Id, student.Name, student.Bio, student.Email)
}

func checkUserId(user User) bool {
	if user.Id < 1 && user.Id > 1000 {
		return false
	}
	return true
}

func checkUserName(user User) bool {
	if len(user.Name) < 2 && len(user.Name) > 20 {
		return false
	}
	return true
}
