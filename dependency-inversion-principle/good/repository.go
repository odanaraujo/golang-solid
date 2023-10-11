package good

type UserRepository interface {
	InsertUser(id string, name string) *User
	GetUserByID(user_id string) *User
}
