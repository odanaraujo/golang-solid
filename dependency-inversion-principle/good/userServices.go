package good

var (
	inMemoUsers map[string]User
)

func init() {
	inMemoUsers = make(map[string]User)
}

type UserServices struct {
	Repository UserRepository
}

func (u *UserServices) InsertUser(user_id string, name string) *User {
	return u.Repository.InsertUser(user_id, name)
}

func (u *UserServices) GetUserByID(user_id string) *User {
	return u.Repository.GetUserByID(user_id)
}
