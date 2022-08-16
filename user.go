package go_cms

import "fmt"

type User struct {
	Name    string
	Age     int
	Gender  string
	Address string
}

func Users(user User) User {
	return user
}

func main() {
	var user User
	user.Name = "Pramudya"
	user.Age = 16
	user.Gender = "Male"
	user.Address = "caringin"

	fmt.Println(user)
}
