package bad

import (
	"math/rand"
	"strconv"
)

var (
	inMemoUsers map[string]User
)

func init() {
	inMemoUsers = make(map[string]User)
}

type User struct {
	ID   string
	Name string
}

// baixo nível
func InsertUserIntoMySQLRepo(user_id string, name string) *User {
	user := User{
		ID:   user_id,
		Name: name,
	}
	inMemoUsers[user_id] = user
	return &user
}

// baixo nível
func GetUserByIDFromMySQLRepo(user_id string) *User {
	user, ok := inMemoUsers[user_id]
	if !ok {
		return nil
	}

	return &user
}

// Alto nível
func GetUserByIDServices(user_id string) *User {
	return GetUserByIDFromMySQLRepo(user_id)
}

// Alto nível
func InsertUserServices(name string) *User {
	user_id := strconv.Itoa(rand.Intn(100000))

	return InsertUserIntoMySQLRepo(user_id, name)
}
